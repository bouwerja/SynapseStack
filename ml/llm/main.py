from llama_cpp import Llama

llm = Llama.from_pretrained(
    repo_id = "hugging-quants/Llama-3.2-3B-Instruct-Q4_K_M-GGUF",
    filename = "llama-3.2-3b-instruct-q4_k_m.gguf",
)

llm.create_chat_completion(
    message = {
        "role" : "user",
        "content" : "What is the capital of France?"
    }
)
