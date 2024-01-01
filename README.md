## Go tool to generate ascii Image

#### Usage
build the binary and use it with supported flags 
Make sure to Zoom out full screen terminal to see the art 
For output .txt files, reduce the font size in any editor to see the art   

#### Supported flags
-path  (Path for the image) (Jpeg,Jpeg,Png are supported types)  (Required flag)
-method (method for pixel-to-brightness conversion) (Average,Lightness,Luminosity are available conversion methods) (Optional flag)
-save (To save the ascii in .txt file) (Optional flag)
-h (Help flag)
 
Example: 
Run 
```bash
 go build -o Ascii-generator 
 ./Ascii-generator -path path_to_your_image -method method_for_brightness
```

For help 
Run 
```bash
./Ascii-generator -h
```

### Future Scope
 
Turning this thing to a cli-based app, with user preferred brightness configuration 
