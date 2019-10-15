package graphparser_test

import (
	"os"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "./" // graphparser
)

func fopen(filePath string) (map[string]*Graph, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	g, err := New(file)
	if err != nil {
		return nil, err
	}
	return g, nil
}

var _ = Describe("Graphparser", func() {
	Describe("Instantiation", func() {
		Context("Load Graphml file", func() {
			It("Graphparser", func() {
				g, err := fopen("../testdata/test01.graphml")
				Expect(err).To(BeNil())
				Expect(g).NotTo(BeNil())
				// グラフの接続関係のテスト
				Expect(len(g)).To(Equal(1))       // グラフ数は1
				Expect(len(g["G"].Nodes)).To(Equal(2)) // ノード数は2
				Expect(len(g["G"].Edges)).To(Equal(1)) // エッジ数は1
			})
			It("Parser Error", func() {
				g, err := fopen("../testdata/test_fail00.graphml")
				Expect(err).NotTo(BeNil())
				Expect(g).To(BeNil())
			})
		})
		Context("Load Graphml files", func() {
			It("Simple.graphml", func() {
				g, err := fopen("../testdata/simple.graphml")
				Expect(err).To(BeNil())
				Expect(g).NotTo(BeNil())
				// _, err = GraphML(g)
				// Expect(err).To(BeNil())
				// グラフの接続関係のテスト
				Expect(len(g)).To(Equal(1))       // グラフ数は1
				Expect(len(g["G"].Nodes)).To(Equal(11)) // ノード数は11
				Expect(len(g["G"].Edges)).To(Equal(12)) // エッジ数は12
			})
		})

	})
})
