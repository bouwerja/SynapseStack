# Llama 3.2 3B Q4_K_M

[HuggingFace](https://huggingface.co/hugging-quants/Llama-3.2-3B-Instruct-Q4_K_M-GGUF)

---

## Installation

1. Use with llama.cpp
   `brew install llama.cpp`

2. Invoke llama.cpp server

```bash
llama-server --hf-repo hugging-quants/Llama-3.2-3B-Instruct-Q4_K_M-GGUF --hf-file llama-3.2-3b-instruct-q4_k_m.gguf -c 2048
```

- Setup

Clone llama.cpp

```git
git clone https://github.com/ggerganov/llama.cpp
```

Move into llama.cpp folder

```bash
cd llama.cpp && LLAMA_CURL=1 make
```

Run interface

```bash
./llama-cli --hf-repo hugging-quants/Llama-3.2-3B-Instruct-Q4_K_M-GGUF --hf-file llama-3.2-3b-instruct-q4_k_m.gguf -p "The meaning to life and the universe is"
```

---

## Python

```bash
pip install llama-cpp-python
```
