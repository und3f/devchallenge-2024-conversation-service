#!/usr/bin/env python3

from pysentimiento import create_analyzer
import torch
from transformers import AutoModelForCausalLM, AutoTokenizer

create_analyzer(task="sentiment", lang="en")

AutoModelForCausalLM.from_pretrained("numind/NuExtract-tiny", torch_dtype=torch.bfloat16, trust_remote_code=True)
AutoTokenizer.from_pretrained("numind/NuExtract-tiny", trust_remote_code=True)
