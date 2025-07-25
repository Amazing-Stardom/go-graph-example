package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/dominikbraun/graph"
)

// FamilyMember represents a person in our dataset. The ID is the unique identifier.
type FamilyMember struct {
	ID         string
	Name       string
	ParentName string
}

// Edge represents a directed connection from a parent to a child.
type Edge struct {
	From string // Parent's ID
	To   string // Child's ID
}

// GraphOutput is a simple representation of the graph's structure,
// used for visualization and other operations.
type GraphOutput struct {
	Nodes []FamilyMember
	Edges []Edge
}

// 1. --- Building the Graph ---
// buildFamilyGraph takes raw data and constructs a graph representation.
func buildFamilyGraph(members []FamilyMember) (*GraphOutput, graph.Graph[string, FamilyMember]) {
	// Use the member's ID as the unique hash for the graph.
	g := graph.New(func(member FamilyMember) string {
		return member.ID
	}, graph.Directed())

	// A map to easily find a person's ID by their name.
	nameToID := make(map[string]string)

	// Step 1: Add all family members as nodes (vertices) to the graph.
	for _, member := range members {
		g.AddVertex(member)
		nameToID[member.Name] = member.ID
	}

	var edges []Edge
	// Step 2: Add the parent-child relationships as edges.
	for _, member := range members {
		if member.ParentName != "" {
			if parentID, ok := nameToID[member.ParentName]; ok {
				// Add an edge from the parent to the child.
				g.AddEdge(parentID, member.ID)
				edges = append(edges, Edge{From: parentID, To: member.ID})
			}
		}
	}

	graphOut := &GraphOutput{
		Nodes: members,
		Edges: edges,
	}

	return graphOut, g
}

// 2. --- Visualizing the Graph ---
// renderPNG generates a PNG image from the graph structure using Graphviz.
// It creates a file named `family_tree.png`.
func renderPNG(graphOut *GraphOutput, outPath string) error {
	var b strings.Builder
	b.WriteString("digraph G {\n")
	b.WriteString("  rankdir=TB;\n") // Top-to-Bottom layout for a family tree
	b.WriteString("  node [shape=box, style=rounded];\n")

	// Add nodes to the dot file
	for _, n := range graphOut.Nodes {
		label := n.Name // Use the person's name as the label
		b.WriteString(fmt.Sprintf("  \"%s\" [label=\"%s\"];\n", n.ID, label))
	}

	// Add edges to the dot file
	for _, e := range graphOut.Edges {
		b.WriteString(fmt.Sprintf("  \"%s\" -> \"%s\";\n", e.From, e.To))
	}
	b.WriteString("}\n")

	// Execute the 'dot' command to render the PNG
	cmd := exec.Command("dot", "-Tpng", "-o", outPath)
	cmd.Stdin = strings.NewReader(b.String())

	fmt.Printf("Attempting to generate graph image at '%s'...\n", outPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to generate graph image. Is Graphviz installed?: %w", err)
	}

	fmt.Printf("Successfully generated %s!\n", outPath)
	return nil
}

// 3. --- Performing Graph Operations ---
func main() {
	// --- Sample Data ---
	// A scattered list of people. The graph will organize them.
	familyData := []FamilyMember{
		{ID: "1", Name: "Jordon D", ParentName: ""},
		{ID: "2", Name: "Robert Smith", ParentName: ""},
		{ID: "3", Name: "Danis Jordan", ParentName: "Jordon D"},
		{ID: "4", Name: "John Smith", ParentName: "Robert Smith"},
		{ID: "5", Name: "Maria Smith", ParentName: "Robert Smith"},
		{ID: "6", Name: "Leo Smith", ParentName: "John Smith"},
	}

	// --- Build and Visualize ---
	graphOutput, familyGraph := buildFamilyGraph(familyData)
	if err := renderPNG(graphOutput, "family_tree.png"); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n--- Graph Operations ---")

	// --- Operation 1: Depth-First Search (DFS) ---
	// Traces a line of descent from an ancestor.
	fmt.Println("\n[DFS] Tracing all descendants of Robert Smith:")
	_ = graph.DFS(familyGraph, "2", func(value string) bool {
		member, _ := familyGraph.Vertex(value)
		fmt.Printf(" -> %s\n", member.Name)
		return false // Continue traversal
	})

	// --- Operation 2: Breadth-First Search (BFS) ---
	// Explores the tree level by level (generation by generation).
	fmt.Println("\n[BFS] Exploring family by generation starting from Jordon D:")
	_ = graph.BFS(familyGraph, "1", func(value string) bool {
		member, _ := familyGraph.Vertex(value)
		fmt.Printf(" - Found: %s\n", member.Name)
		return false // Continue traversal
	})
	
	// --- Operation 3: Find Roots (Original Ancestors) ---
	// Finds all nodes with no incoming edges (no parents).
	fmt.Println("\n[Roots] Finding the original ancestors:")
	for _, member := range familyData {
		// PredecessorMap gives us all parents of a node.
		predecessors, _ := familyGraph.PredecessorMap()
		if len(predecessors[member.ID]) == 0 {
			fmt.Printf(" - %s is a root ancestor.\n", member.Name)
		}
	}

	// --- Operation 4: Find Leaves (Youngest Generation) ---
	// Finds all nodes with no outgoing edges (no children).
	fmt.Println("\n[Leaves] Finding members with no children:")
	for _, member := range familyData {
		// AdjacencyMap gives us all children of a node.
		successors, _ := familyGraph.AdjacencyMap()
		if len(successors[member.ID]) == 0 {
			fmt.Printf(" - %s has no children.\n", member.Name)
		}
	}
}
