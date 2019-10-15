package graphparser

import(
	// "fmt"
	"io"
	"strconv"
	"io/ioutil"
	"encoding/xml"
)

type GraphML struct {
	// XMLName     xml.Name  `xml:"graphml"`
	// Xml         string    `xml:",innerxml"`
	// Id     string    `xml:"id,attr"`
	Graphs []xmlGraph `xml:"graph"`
}

type xmlGraph struct {
	Id string    `xml:"id,attr"`
	EdgeDefault string `xml:"edgedefault,attr"`
	Nodes []xmlNode `xml:"node"`
	Edges []xmlEdge `xml:"edge"`
}

type xmlNode struct {
	Id string `xml:"id,attr"`
	Xml         string    `xml:",innerxml"`
}

type xmlEdge struct {
	Id string     `xml:"id,attr"`
	Source string `xml:"source,attr"`
	Target string `xml:"target,attr"`
	Xml         string    `xml:",innerxml"`
}

type Graph struct {
	Nodes map[string]*Node
	Edges map[string]*Edge
}

type Node struct {
	Sources []*Edge
	Targets []*Edge
	Xml     string
}

type Edge struct {
	Source *Node
	Target *Node
	Xml     string
}

func New(r io.Reader) (map[string]*Graph, error) {
    xmlByte, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

    data := new(GraphML)
    if err := xml.Unmarshal(xmlByte, data); err != nil {
        return nil, err
    }

	// すべてのEdgeについて，Nodesを紐付け
	gmap := make(map[string]*Graph)
	for _, graph := range data.Graphs {
		g := &Graph{Nodes: make(map[string]*Node), Edges: make(map[string]*Edge)}

		// ノードの登録
		//   接続されているエッジは，エッジの読み込み時に登録する
		for _, node := range graph.Nodes {
			g.Nodes[node.Id] = &Node{Sources: make([]*Edge,0), Targets: make([]*Edge,0), Xml: node.Xml}
		}
		// エッジの登録
		//   Source, Targetに該当するノードがあれば，一緒に登録しておく
		//   EdgeのIDはオプションなので，ついていない場合は記述順に番号を振る
		for eidx, edge := range graph.Edges {
			e :=&Edge{}
			if v, ok := g.Nodes[edge.Source]; ok {
				e.Source  = v
				v.Sources = append(v.Sources, e)
			}
			if v, ok := g.Nodes[edge.Target]; ok {
				e.Target  = v
				v.Targets = append(v.Targets, e)
			}
			eId := edge.Id
			if  eId == "" {
				eId = strconv.Itoa(eidx)
			}
			e.Xml = edge.Xml
			g.Edges[eId] = e
		}
		gmap[graph.Id] = g
	}
	return gmap, nil
}


