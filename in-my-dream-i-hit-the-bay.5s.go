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
	_ "embed"
	"fmt"
	"log"
	"math/rand"

	"github.com/leaanthony/go-ansi-parser"
)

const (
	secondEntryText = "thoughts"
)

var (
	// The .B64 file is obtained by base64-encoding this image:
	// https://img.favpng.com/11/18/10/tightrope-walking-computer-icons-youtube-png-favpng-qcQdAyeKgpbCPdmQeaU9757iD.jpg
	//go:embed man-walking-on-a-tightrope.B64
	manWalkingOnATightropeB64 string

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
	//go:embed "stack-overflow-logo.ans"
	stackOverflowLogo string

	firstEntryAlternatingTexts = []string{"Everything will be fine.", "Nah, it's just a fairy tale."}
)

func mustBeNil(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

func printThumbnail(thumbnail string) {
	fmt.Printf("| templateImage=%s\n", string(thumbnail))
}

func printFirstEntry(texts []string) {
	idx := rand.Intn(len(texts))
	fmt.Println(texts[idx])
}

func printSecondEntry(text, img string) {
	fmt.Println(text)

	start := 0
	for i := 0; i < len(img); i++ {
		if img[i] == '\n' {
			// The ansi library is used because
			// it reset all modes (i.e., ESC[0m) at the end of each ANSI sequence.
			// The implication is that when printing '--',
			// those strings won't be stylized by the last ANSI sequence
			// as the original .ans file does not necessarily reset all modes at the end of each sequence.
			row, err := ansi.Parse(img[start:i])
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
	printThumbnail(manWalkingOnATightropeB64)
	printSeparator()
	printFirstEntry(firstEntryAlternatingTexts)
	printSeparator()
	printSecondEntry(secondEntryText, stackOverflowLogo)
}
