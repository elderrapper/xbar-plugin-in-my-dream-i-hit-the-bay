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
	"time"

	"github.com/leaanthony/go-ansi-parser"
)

const (
	firstEntryText1 = "Everything will be fine."
	firstEntryText2 = "Nah, it's just a fairy tale."
	secondEntryText = "thoughts"
	monoSpaceFont   = "font=Menlo"

	esc                 = "\x1B"
	moveUpTmpl          = esc + "[%dA"
	eraseCurrentLineSeq = esc + "[2K"
)

var (
	// To get the .B64 file, follow these steps:
	// 1. Download https://img.favpng.com/11/18/10/tightrope-walking-computer-icons-youtube-png-favpng-qcQdAyeKgpbCPdmQeaU9757iD.jpg
	// 2. Resize the image to be 30x19 via https://www.iloveimg.com/resize-image#resize-options,pixels
	//    because a menu bar icon cannot be too large.
	// 3. Base64-encode the resulting image via https://elmah.io/tools/base64-image-encoder/.
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
	// 3. Use some tool to crop part of the top and bottom blanks of the image.
	//    Otherwise, those blanks become empty lines in the menu,
	//    and we don't want too much of them as screen space is precious.
	// 4. Widen the image via https://www.iloveimg.com/resize-image#resize-options,pixels
	//    because the spacing between the lines in the menu makes the image look vertically stretched.
	// 5. Convert the resulting file to be an .ans one (width = 100) using
	//    https://manytools.org/hacker-tools/convert-images-to-ascii-art/.
	// 6. Manually remove the first ANSI sequence (i.e., ESC[30;107m) to prevent the parsing error from xbar
	//    because it's still using leaanthony/go-ansi-parser@v1.2.0
	//    (https://github.com/matryer/xbar/blob/c6fa2be71000f6665e2b68011506d4c0dce24268/app/go.mod#L12),
	//    which does not include the fix for https://github.com/leaanthony/go-ansi-parser/issues/3.
	//go:embed "stack-overflow-logo.ans"
	stackOverflowLogo string
)

func mustBeNil(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

func printThumbnail(thumbnail string) {
	fmt.Printf("| templateImage=%s\n", thumbnail)
}

// printSecondEntry returns the height of the image.
func printSecondEntry(text, image string) int {
	fmt.Println(text)

	start, n := 0, 0
	for i := 0; i < len(image); i++ {
		if image[i] == '\n' {
			// The ansi library is used because
			// it resets all modes (i.e., ESC[0m) at the end of each ANSI sequence.
			// The implication is that when printing '--',
			// those strings won't be stylized by the last ANSI sequence
			// as the original .ans file does not necessarily reset all the modes at the end of each sequence.
			row, err := ansi.Parse(image[start:i])
			mustBeNil(err, "failed to parse the row into ANSI sequences")

			fmt.Printf("--%s | %s\n", ansi.String(row), monoSpaceFont)
			start = i + 1
			n++
		}
	}
	return n
}

func printSeparator() {
	fmt.Println("---")
}

func moveCursorUp(n int) {
	fmt.Printf(moveUpTmpl, n)
}

func eraseCurrentLine() {
	fmt.Print(eraseCurrentLineSeq)
}

func main() {
	printThumbnail(manWalkingOnATightropeB64)
	printSeparator()
	fmt.Println(firstEntryText1)
	printSeparator()
	n := printSecondEntry(secondEntryText, stackOverflowLogo)

	time.Sleep(2 * time.Second)
	// 3 = the first entry + the separator between the first entry and the second entry + text portion of the second entry
	moveCursorUp(n + 3)
	eraseCurrentLine()
	fmt.Println(firstEntryText2)
}
