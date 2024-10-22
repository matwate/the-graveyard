# Calendar Generator

This project is a Python application that generates a rotation-based calendar for a given year and a list of names. The generated calendar is saved as an Excel file.

## Files

- `calendario.py`: This is the main script that generates the calendar. It contains the `createCronogram` function which takes a year and a list of names as input and generates the calendar.

- `Ui.py`: This script creates a simple GUI for the application. It allows the user to input the year and the names, and then generates the calendar when the "Generar" button is clicked.

- `excel/result.xlsx`: This is where the generated calendar is saved.

## Usage

1. Run `Ui.py` to start the application.
2. Enter the names in the "Nombre" fields.
3. Enter the year in the "Año" field.
4. Click the "Generar" button to generate the calendar. The calendar will be saved as `excel/result.xlsx`.

## Requirements

- Python 3
- openpyxl
- customtkinter

## Installation

1. Clone this repository.
2. Install the required Python packages using pip:

```sh
pip install openpyxl customtkinter
```
3. Run `ui.py` to start the application

