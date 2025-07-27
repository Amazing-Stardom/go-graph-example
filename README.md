# Go Graph Example: Family Tree Traversal

This project demonstrates how to model and query relational data using a graph data structure in Go. It uses the `dominikbraun/graph` library to build a family tree, visualize it, and perform common graph traversal operations like Depth-First Search (DFS) and Breadth-First Search (BFS).

This code is the official implementation for the blog post: **Understanding Graph Implementation and Query Operations in Go**.

## Features

*   **Graph Construction**: Builds a directed acyclic graph (DAG) from a simple list of family members.
*   **Graph Visualization**: Generates a `.png` image of the family tree using Graphviz for clear visualization.
*   **Depth-First Search (DFS)**: Traverses the graph to find all descendants of a specific ancestor.
*   **Breadth-First Search (BFS)**: Explores the family tree level by level, showing each generation.
*   **Root and Leaf Node Identification**: Finds the original ancestors (root nodes) and the youngest generation (leaf nodes) in the tree.

## Prerequisites

Before you can run this project, you need to have the following installed:

*   **Go**: The programming language used for this project.
*   **Graphviz**: A graph visualization software used to render the family tree image. Make sure the `dot` command is available in your system's PATH.

## How to Run

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/Amazing-Stardom/go-graph-example.git
    cd go-graph-example
    ```

2.  **Install dependencies:**
    This project uses `go mod` for dependency management. Run the following command to download the necessary libraries:
    ```sh
    go mod tidy
    ```

3.  **Run the program:**
    Execute the main Go file to build the graph, generate the image, and run the queries:
    ```sh
    go run main.go
    ```

## Expected Output

After running the program, you will see the following:

1.  A new file named **`family_tree.png`** will be generated in the project directory, showing the visual representation of the family tree.
2.  The console will display the output from the different graph operations:

    ```
    Attempting to generate graph image at 'family_tree.png'...
    Successfully generated family_tree.png!

    --- Graph Operations ---

    [DFS] Tracing all descendants of Robert Smith:
     -> Robert Smith
     -> Maria Smith
     -> John Smith
     -> Leo Smith

    [BFS] Exploring family by generation starting from Jordon D:
     - Found: Jordon D
     - Found: Danis Jordan

    [Roots] Finding the original ancestors:
     - Jordon D is a root ancestor.
     - Robert Smith is a root ancestor.

    [Leaves] Finding members with no children:
     - Danis Jordan has no children.
     - Maria Smith has no children.
     - Leo Smith has no children.
    ```
