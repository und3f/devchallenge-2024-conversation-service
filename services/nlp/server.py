#!/usr/bin/env python3

import logging
from pysentimiento import create_analyzer
from transformers import AutoModelForCausalLM, AutoTokenizer
import torch
from flask import Flask, request, jsonify
import json

def predict_NuExtract(model, tokenizer, text, schema, example=["", "", ""]):
    schema = json.dumps(json.loads(schema), indent=4)
    input_llm =  "<|input|>\n### Template:\n" +  schema + "\n"
    for i in example:
      if i != "":
          input_llm += "### Example:\n"+ json.dumps(json.loads(i), indent=4)+"\n"
    
    input_llm +=  "### Text:\n"+text +"\n<|output|>\n"
    input_ids = tokenizer(input_llm, return_tensors="pt",truncation = True, max_length=4000).to("cpu")

    output = tokenizer.decode(model.generate(**input_ids)[0], skip_special_tokens=True)
    return output.split("<|output|>")[1].split("<|end-output|>")[0]

model = AutoModelForCausalLM.from_pretrained("numind/NuExtract-tiny", torch_dtype=torch.bfloat16, trust_remote_code=True)
tokenizer = AutoTokenizer.from_pretrained("numind/NuExtract-tiny", trust_remote_code=True)

model.to("cpu")
model.eval()

analyzer = create_analyzer(task="sentiment", lang="en")

FORMAT = '%(asctime)s %(message)s'
logging.basicConfig(format=FORMAT, level=logging.INFO)
logger = logging.getLogger(__name__)

app = Flask(__name__)

@app.route('/health', methods=['GET'])
def health():
    return jsonify({"status": "ok"})

@app.route('/emotion', methods=['POST'])
def predict_emotion():
    logger.info('/emotion')

    text = request.get_json()["text"]

    out = analyzer.predict(text)
    resp = {
        "output": out.output,
        "probas": out.probas
    }
    return jsonify(resp)

@app.route('/extract', methods=['POST'])
def predict_extract():
    logger.info('/extract')

    json = request.get_json()
    text = json["text"]
    schema = json["schema"]

    prediction = predict_NuExtract(model, tokenizer, text, schema, example=["","",""])
    return prediction

if __name__ == '__main__':
    from waitress import serve
    serve(app, host="0", port=8080)
