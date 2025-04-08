import random

def random_go_code():
    # Generate the start of the Go code
    start = 'Go'
    if random.random() > 0.5:
        start += ', '
    end = '!'
    
    # Add some Go-specific characters
    codes = [
        "go",
        ", ",
        "."
    ]
    return f"{start} {random.choice(codes)} {end}!"

# Call the function and print the result
print(random_go_code())
