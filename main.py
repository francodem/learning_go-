input = "Olleh"

word = "Hello"

def analizer(input, word):
    if word == input[::-1]:
        return True
    else:
        return False


analizer(input, word)