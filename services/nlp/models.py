from pysentimiento import create_analyzer
from transformers import AutoModelForCausalLM, AutoTokenizer
import torch

model = AutoModelForCausalLM.from_pretrained("numind/NuExtract-tiny", torch_dtype=torch.bfloat16, trust_remote_code=True)
tokenizer = AutoTokenizer.from_pretrained("numind/NuExtract-tiny", trust_remote_code=True)

model.to("cpu")
model.eval()

analyzer = create_analyzer(task="sentiment", lang="en")

