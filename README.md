# Code Plagiarism Checker

## Overview

This project is a Code Plagiarism Checker designed to identify and highlight similarities between source code files. It serves as a tool for developers, educators, and organizations to ensure the integrity of code submissions and detect potential instances of plagiarism.

## Features

- **File Comparison:** Utilizes advanced algorithms to compare source code files and identify similarities.
- **Configurable Thresholds:** Allows users to set similarity thresholds based on their specific requirements.
- **Multi-Language Support:** Supports various programming languages for comprehensive plagiarism detection.
- **User-Friendly Interface:** Provides an easy-to-use command-line interface for efficient code analysis.

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/code-plagiarism-checker.git
    ```

2. Install dependencies:

    ```bash
    cd code-plagiarism-checker
    pip install -r requirements.txt
    ```

## Usage

1. Navigate to the project directory:

    ```bash
    cd code-plagiarism-checker
    ```

2. Run the code plagiarism checker:

    ```bash
    python plagiarism_checker.py -d /path/to/code/directory
    ```

    Replace `/path/to/code/directory` with the path to the directory containing the source code files you want to analyze.

3. Adjust settings as needed:

    ```bash
    python plagiarism_checker.py -d /path/to/code/directory -t 0.85 -l python,java
    ```

    - `-t` sets the similarity threshold (default is 0.80).
    - `-l` specifies the programming languages to consider (default is all supported languages).


## Contributing

Contributions are welcome! If you find a bug or have an enhancement in mind, please open an issue or submit a pull request.

## License

This Code Plagiarism Checker is open-source software licensed under the [MIT License](LICENSE).

## Acknowledgments

- Special thanks to [List of Algorithms](https://github.com/TheAlgorithms) for providing foundational algorithms for code comparison.

## Disclaimer

This tool is designed to assist in identifying potential instances of plagiarism and should not be the sole basis for any decisions. Always consider additional context and evidence when evaluating code submissions.
