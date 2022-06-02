//usr/bin/env go run $0 $@; exit
package main

// <xbar.title>In My Dream I Hit The Bay</xbar.title>
// <xbar.version>v1.0</xbar.version>
// <xbar.author>Hsing-Yu Chen</xbar.author>
// <xbar.author.github>davidhsingyuchen</xbar.author.github>
// <xbar.desc>This plugin is used to produce the MV of In My Dream I Hit The Bay.</xbar.desc>
// <xbar.image>https://img.favpng.com/11/18/10/tightrope-walking-computer-icons-youtube-png-favpng-qcQdAyeKgpbCPdmQeaU9757iD.jpg</xbar.image>
// <xbar.dependencies>go</xbar.dependencies>
// <xbar.abouturl>https://github.com/davidhsingyuchen/xbar-plugin-in-my-dream-i-hit-the-bay</xbar.abouturl>

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/leaanthony/go-ansi-parser"
)

const (
	// The .B64 file is obtained by base64-encoding this image:
	// https://img.favpng.com/11/18/10/tightrope-walking-computer-icons-youtube-png-favpng-qcQdAyeKgpbCPdmQeaU9757iD.jpg
	manWalkingOnATightropeB64Path = `man-walking-on-a-tightrope.B64`
	// To get the .ans file, follow these steps:
	// 1. Download https://commons.wikimedia.org/wiki/File:Stack_Overflow_icon.svg.
	//    The file extension will be png instead of svg
	//    if you simply right click at the image and click "Save image as...".
	// 2. Convert the image file to be a .jpg one.
	//    Otherwise, when the image is converted to be an .ans one,
	//    the color of a "transparent" pixel in the .png file will be left blank (e.g., ESC[38;5;m)
	//    and the ANSI parser will complain about that.
	// 3. Convert th e.jpg file to be an .ans one using
	//    https://manytools.org/hacker-tools/convert-images-to-ascii-art/.
	stackOverflowLogoPath = "stack-overflow-logo.ans"
	secondEntryText       = "thoughts"
)

var firstEntryAlternatingTexts = []string{"Everything will be fine.", "Nah, it's just a fairy tale."}

func mustBeNil(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

func printThumbnail(path string) {
	thumbnail, err := os.ReadFile(path)
	mustBeNil(err, "failed to read the thumbnail from the file system")
	fmt.Printf("| templateImage=%s\n", string(thumbnail))
}

func printFirstEntry(texts []string) {
	idx := rand.Intn(len(texts))
	fmt.Println(texts[idx])
}

func printSecondEntry(text, imgPath string) {
	fmt.Println(text)

	img, err := os.ReadFile(imgPath)
	mustBeNil(err, "failed to read the image for the second entry from the file system")

	str := string(img)
	start := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '\n' {
			// The ansi library is used because
			// it reset all modes (i.e., ESC[0m) at the end of each ANSI sequence.
			// The implication is that when printing '--',
			// those strings won't be stylized by the last ANSI sequence
			// as the original .ans file does not necessarily reset all modes at the end of each sequence.
			row, err := ansi.Parse(str[start:i])
			mustBeNil(err, "failed to parse the row into ANSI sequences")

			fmt.Printf("--%s\n", ansi.String(row))
			start = i + 1
		}
	}
}

func printSeparator() {
	fmt.Println("---")
}

func main() {
	printThumbnail(manWalkingOnATightropeB64Path)
	printSeparator()
	printFirstEntry(firstEntryAlternatingTexts)
	printSeparator()
	printSecondEntry(secondEntryText, stackOverflowLogoPath)
}
