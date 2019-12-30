$(() => {
    let toolbarOptions = {
        container: [
            ['bold', 'italic', 'underline', 'strike'],        // toggled buttons
            ['blockquote', 'code-block'],

            [{'header': 1}, {'header': 2}],               // custom button values
            [{'list': 'ordered'}, {'list': 'bullet'}],
            [{'script': 'sub'}, {'script': 'super'}],      // superscript/subscript
            [{'indent': '-1'}, {'indent': '+1'}],          // outdent/indent
            [{'direction': 'rtl'}],                         // text direction

            [{'size': ['small', false, 'large', 'huge']}],  // custom dropdown
            [{'header': [1, 2, 3, 4, 5, 6, false]}],
            ['link', 'image', 'video', 'formula'],          // add's image support
            [{'color': []}, {'background': []}],          // dropdown with defaults from theme
            [{'font': []}],
            [{'align': []}],

            ['clean']                                         // remove formatting button
        ]
    };
    let quill = new Quill('#editor', {
        modules: {
            toolbar: toolbarOptions
        },
        theme: 'snow'
    });
    quill.getModule('toolbar').addHandler('image', () => {
        selectLocalImage();
    });

    /**
     * Step1. select local image
     *
     */
    function selectLocalImage() {
        const input = document.createElement('input');
        input.setAttribute('type', 'file');
        input.click();

        // Listen upload local image and save to server
        input.onchange = () => {
            const file = input.files[0];

            // file type is only image.
            if (/^image\//.test(file.type)) {
                saveToServer(file);
            } else {
                console.warn('You could only upload images.');
            }
        };
    }

    /**
     * Step2. save to server
     *
     * @param {File} file
     */
    function saveToServer(file) {
        const fd = new FormData();
        fd.append('image', file);

        const xhr = new XMLHttpRequest();
        xhr.open('POST', '/api/nopiser/picture', true);
        xhr.onload = () => {
            if (xhr.status === 200) {
                // this is callback data: url
                const url = JSON.parse(xhr.responseText).url;
                insertToEditor(url);
            }
        };
        xhr.send(fd);
    }

    /**
     * Step3. insert image url to rich editor.
     *
     * @param {string} url
     */
    function insertToEditor(url) {
        // push image url to rich editor.
        const range = quill.getSelection();
        quill.insertEmbed(range.index, 'image', `${url}`);
    }

// quill editor add image handler

    $("#submit").on("click", () => {
        let delta = quill.getContents();
        $.post("/api/ritin/article", JSON.stringify({content: JSON.stringify(delta.ops)}), (data) => {
            window.location.href = "/api/ritin/article/" + data.articleId
        })
    });
});
