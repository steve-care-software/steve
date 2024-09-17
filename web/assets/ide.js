function adjustHeight(className) {
    var elements = document.getElementsByClassName(className);
    for (i = 0; i < elements.length; i++) {
        elements[i].onkeydown = function(e) {
            if (e.keyCode === 9) {
                if (e.keyCode === 9) {
                    // get the current cursor position
                    const position = this.selectionStart;
            
                    // get the text before and after the cursor position
                    const before = this.value.substring(0, position);
                    const after = this.value.substring(position, this.value.length);
            
                    // insert the new text at the cursor position
                    var text = "    ";
                    this.value = before + text + after;
            
                    // set the cursor position to after the newly inserted text
                    this.selectionStart = this.selectionEnd = position + text.length;
                    return false;
                }
            }
        }

        elements[i].onkeyup = function(e) {
            // auto-resize when there is more content
            this.style.height = "5px";
            this.style.height = (this.scrollHeight)+"px";
        }
        
         // auto-resize when there is more content at load:
         elements[i].style.height = "5px";
         elements[i].style.height = (elements[i].scrollHeight)+"px";
    }
};

adjustHeight("ide");