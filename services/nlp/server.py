#!/usr/bin/env python3

import logging
from flask import Flask, request, jsonify

import nuextract
import models

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

    out = models.analyzer.predict(text)
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

    prediction = nuextract.predict_NuExtract(models.model, models.tokenizer, text, schema, example=["","",""])
    return prediction

if __name__ == '__main__':
    from waitress import serve
    serve(app, host="0", port=8082)
