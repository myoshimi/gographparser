package graphparser

import(
	"fmt"
	"encoding/xml"
)

type GraphML struct {
	// XMLName     xml.Name  `xml:"graphml"`
	// Xml         string    `xml:",innerxml"`
	Id          string    `xml:"id,attr"`
	Graphs []Graph `xml:"graph"`
}

type Graph struct {
	Id string    `xml:"id,attr"`
	Nodes []Node `xml:"node"`
	Edges []Edge `xml:"edge"`
}

type Node struct {
	Id string `xml:"id,attr"`
}

type Edge struct {
	Id string     `xml:"id,attr"`
	Source string `xml:"source,attr"`
	Target string `xml:"target,attr"`
}

func New() (*GraphML, error) {
    xmlStr := `
<?xml version="1.0" encoding="UTF-8"?>

<graphml xmlns="http://graphml.graphdrawing.org/xmlns"
xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
xsi:schemaLocation="http://graphml.graphdrawing.org/xmlns
http://graphml.graphdrawing.org/xmlns/1.0/graphml.xsd">

<graph id="G" edgedefault="undirected">
<node id="n0"/>
<node id="n1"/>
<edge id="e1" source="n0" target="n1"/>
</graph>

</graphml>
`
    data := new(GraphML)
    if err := xml.Unmarshal([]byte(xmlStr), data); err != nil {
        fmt.Println("XML Unmarshal error:", err)
        return nil, err
    }
	return data, nil
}
