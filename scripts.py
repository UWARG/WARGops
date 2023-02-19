import os
import subprocess
import shlex

scripts_folder = "./scripts/"
script_files = [filename for filename in os.listdir(scripts_folder) if filename.endswith(".sh")]


print("Available scripts:")
for i, script_file in enumerate(script_files):
    print(f"{i + 1}. {script_file}")



script_num = input("Enter the number     of the script to execute: ")
while not script_num.isdigit() or int(script_num) < 1 or int(script_num) > len(script_files):
    print("Invalid input. Please enter a number between 1 and", len(script_files))
    script_num = input("Enter the number of the script to execute: ")
script_file = script_files[int(script_num) - 1]
script_path = os.path.join(scripts_folder, script_file)

script_args = input(f"Enter any arguments for {script_file} (leave blank for no arguments): ")


print("Executing", script_path, "with arguments", script_args)
# Execute the chosen script with the given arguments using subprocess
if script_args.strip() == "":
    output = subprocess.check_output(["sh", script_path], stderr=subprocess.STDOUT)
else:
    output = subprocess.check_output(["sh"] + shlex.split(script_path + " " + script_args), shell=True, stderr=subprocess.STDOUT)
print(output.decode())

