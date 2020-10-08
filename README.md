# Phasmocrypt

A decrypter/encrypter for Phasmophobia save files.  **PLEASE** only use this in private games where all parties have
agreed to its use.  This is *only* for making the game fun for you and friends, got it?

## Using
Copy the binaries encrypt and decrypt (.exe on Windows) to the save directory (or move out the save data.  A backup
is a good idea!).

Run the decrypter.  You can specify --input and --output values, but double-clicking will take a `saveData.txt` file
by default and output `saveDataDec.txt`.

Make your edits to `saveDataDec.txt` and save the file as `saveDataEdit.txt` (by default).

Run the encrypter.  This takes --input and --output as well, but double-clicking will default to taking `saveDataEdit.txt`
and outputting `saveDataReEnc.txt`.

Copy the output file `saveDataReEnc.txt` (or whatever you named it) to `saveData.txt`.  You may need to turn OFF 
`read-only` permission on the file.  It may also mention uploading the save file to Steam Save Cloud.  Upload local to
cloud.