#!/usr/bin/env python3

from pysentimiento import create_analyzer
from flask import Flask, request, jsonify
import json

analyzer = create_analyzer(task="sentiment", lang="en")

app = Flask(__name__)
@app.route('/emotion', methods=['POST'])
def predict():
    text = request.get_json()["text"]
    print("Analyzing", text)

    out = analyzer.predict(text)
    print(out)

    resp = {
        "output": out.output,
        "probas": out.probas
    }
    return jsonify(resp)

if __name__ == '__main__':
    from waitress import serve
    serve(app, host="0", port=8080)
