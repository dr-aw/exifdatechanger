# Photo Date Update Tool

## Overview

This tool updates the `Date Taken` attribute of photos to the `Date of Shot` (from the EXIF metadata). Simply place the executable file in the directory with your photos, and it will automatically update the dates for you.

## How to Use

1. **Download** or **build** the executable file (`.exe`).

2. **Place** the `.exe` file in the directory containing your photos.

3. **Run** the executable file. It will process all supported image files (`.jpeg`, `.jpg`, `.png`) in the directory.

4. **Wait** for the process to complete. The tool will update the `Date Taken` attribute to the `Date of Shot` from the EXIF data of each photo.

## Supported Formats

- `.jpeg`
- `.jpg`
- `.png`

## Requirements

- Windows OS (`x64` or `arm`, for the `.exe` file)
- Photos with EXIF metadata

## Example

If you have a directory `Photos` with several `.jpeg` and `.png` files, place the executable file inside `Photos` and run it. The tool will update the `Date Taken` attribute of each image based on the EXIF `Date of Shot`.

## License

This tool is provided as-is. No warranty is provided for its use. Feel free to modify and use it according to your needs.