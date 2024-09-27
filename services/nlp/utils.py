import multiprocessing
import torch
import sys

def set_torch_fullCPU():
    num_threads = multiprocessing.cpu_count()

    torch.set_num_threads(num_threads)
    torch.set_num_interop_threads(num_threads)
