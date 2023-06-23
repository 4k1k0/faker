package lorem

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

func TestLorem(t *testing.T) {
	l := faker.New().Lorem()
	tool.Expect(t, "faker.Lorem", fmt.Sprintf("%T", l))
}

func TestWord(t *testing.T) {
	l := faker.New().Lorem()
	tool.Expect(t, true, len(l.Word()) > 0)
}

func TestWords(t *testing.T) {
	l := faker.New().Lorem()
	tool.Expect(t, 2, len(l.Words(2)))
}

func TestSentence(t *testing.T) {
	l := faker.New().Lorem()
	split := strings.Split(l.Sentence(2), " ")
	tool.Expect(t, 2, len(split))
}

func TestSentences(t *testing.T) {
	l := faker.New().Lorem()
	sentences := l.Sentences(2)
	tool.Expect(t, 2, len(sentences))
}

func TestParagraph(t *testing.T) {
	l := faker.New().Lorem()
	split := strings.Split(l.Paragraph(2), ".")
	tool.Expect(t, 3, len(split))
}

func TestParagraphs(t *testing.T) {
	l := faker.New().Lorem()
	split := l.Paragraphs(2)
	tool.Expect(t, 2, len(split))
}

func TestText(t *testing.T) {
	l := faker.New().Lorem()
	text := l.Text(255)
	tool.Expect(t, true, len(text) <= 255)
}

func TestTextNotEmpty(t *testing.T) {
	l := faker.New().Lorem()
	for i := 1; i < 255; i++ {
		text := l.Text(i)
		tool.Expect(t, true, len(text) != 0)
	}
}

func TestBytes(t *testing.T) {
	l := faker.New().Lorem()
	text := l.Bytes(255)
	tool.Expect(t, true, len(text) <= 255)
}
