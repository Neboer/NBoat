var toolbarOptions = [
    ['bold', 'italic', 'underline', 'strike'],        // toggled buttons
    ['blockquote', 'code-block'],

    [{'header': 1}, {'header': 2}],               // custom button values
    [{'list': 'ordered'}, {'list': 'bullet'}],
    [{'script': 'sub'}, {'script': 'super'}],      // superscript/subscript
    [{'indent': '-1'}, {'indent': '+1'}],          // outdent/indent
    [{'direction': 'rtl'}],                         // text direction

    [{'size': ['small', false, 'large', 'huge']}],  // custom dropdown
    [{'header': [1, 2, 3, 4, 5, 6, false]}],

    [{'color': []}, {'background': []}],          // dropdown with defaults from theme
    [{'font': []}],
    [{'align': []}],

    ['clean']                                         // remove formatting button
];

$(function () {
    var quill = new Quill('#quill-container', {
        modules: {
            toolbar: [
                [{header: [1, 2, false]}],
                ['bold', 'italic', 'underline'],
                ['image', 'code-block']
            ]
        },
        // placeholder: 'Compose an epic...',
        theme: 'snow'  // or 'bubble'
    });
    // replace oss_config_obj as my own oss key
    let client = new OSS(~pre_render_oss_config_obj);

    quill.getModule('toolbar').addHandler('image', () => {
        $('#upload-modal').modal();
        // const range = quill.getSelection();
        // quill.insertEmbed(range.index, 'image', `${"https://dss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/logo/logo_redBlue-ece8ad45d6.png"}`);
    });

    $("#upload-oss").on('click',() => {
        client.put("")
    })

    const template = document.getElementById('tippy-popup');
});