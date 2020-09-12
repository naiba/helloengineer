export default {
    height: window.innerHeight - 50,
    mode: "wysiwyg",
    preview: {
        hljs: {
            lineNumber: true,
            style: "dracula"
        }
    },
    toolbar: [
        "emoji",
        "headings",
        "bold",
        "italic",
        "strike",
        "link",
        "|",
        "list",
        "ordered-list",
        "check",
        "outdent",
        "indent",
        "|",
        "quote",
        "line",
        "code",
        "inline-code",
        "insert-before",
        "insert-after",
        "|",
        "upload",
        "record",
        "table",
        "|",
        "undo",
        "redo",
        "|",
        "fullscreen",
        // "edit-mode",
        {
            name: "more",
            toolbar: [
                "both",
                "code-theme",
                "content-theme",
                "export",
                "outline",
                "preview"
                // "devtools",
                // "info",
                // "help"
            ]
        }
    ]
} as any