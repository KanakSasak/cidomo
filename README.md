# cidomo

## Encrypt and Set Wallpaper README

This Go code encrypts files in a specified directory, generates a wallpaper image, and sets it as the desktop wallpaper. Additionally, it logs the details of encrypted files.

### Prerequisites

- Go programming language installed
- Basic understanding of Go programming concepts
- Windows operating system

### Instructions

1. Modify the `root` variable in the code to specify the root directory where encryption should start. For Windows systems, the default is `C:\`.

2. Update the `key` variable to set the encryption key. The key should be a 16-byte byte array.

3. Run the program using the command `go run main.go` in the terminal.

4. The program will encrypt files with extensions `.docx`, `.xlsx`, and `.pptx` in the specified directory and its subdirectories.

5. The encrypted files will be saved with a `.goblok` extension in the same location as the original files.

6. The program will generate a new wallpaper image with the text "silahkan periksa file anda" and save it as `output.png` in the temporary folder.

7. The generated wallpaper image will be set as the desktop wallpaper.

8. The program will create a log file named `encryption.log` in the same directory. The log file will contain details of each encrypted file in the format `filename with extension | encrypted filename with extension | directory`.

9. After execution, check the log file for the encrypted file details.

### Customization

- To encrypt files with different extensions, modify the `ext` comparison in the code to include the desired file extensions.

- To change the text on the wallpaper image, modify the `addLabel` function in the code.

- To set the wallpaper image from a different location, update the `tempFile` variable with the desired file path.

### Limitations and Caution

- The code provided is for educational purposes and should be used responsibly.

- Encrypting files indiscriminately can have serious consequences. Ensure you have appropriate permissions and take necessary precautions before running the code.

- Changing the desktop wallpaper requires administrative privileges.

- Always back up important files before encryption.

### License

This code is released under the [MIT License](LICENSE).


