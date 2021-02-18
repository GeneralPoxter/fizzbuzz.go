const form = document.getElementById('form');
const output = document.getElementById('output');
const outputText = document.getElementById('output-text');
const outputCopy = document.getElementById('output-copy');

form.addEventListener('submit', async function(event) {
	event.preventDefault();
	let params = new URLSearchParams(new FormData(this)).toString();
	const res = await fetch('/fizzbuzz?' + params).then((res) => res.text());
	let disp = function() {
		outputText.innerText = res;
		output.style.height = 'auto';
		output.style.maxHeight = '500px';
	};
	if (output.style.maxHeight == '0') {
		disp;
	} else {
		output.style.maxHeight = '0';
		setTimeout(disp, 300);
	}
});

function setAttributes(ele, attrs) {
	for (let key in attrs) {
		ele.setAttribute(key, attrs[key]);
	}
}

let conditions = 0;

function addCondition(key = '', str = '') {
	let row = document.createElement('div');
	setAttributes(row, { class: 'row', draggable: 'true', ondragover: 'return false' });

	let keyBoxDiv = document.createElement('div');
	let keyBox = document.createElement('input');
	setAttributes(keyBox, { class: 'text', name: 'cond-key', type: 'number', value: key, in: '1', max: '10000' });
	keyBoxDiv.appendChild(keyBox);

	let strBoxDiv = document.createElement('div');
	let strBox = document.createElement('input');
	setAttributes(strBox, { class: 'text', name: 'cond-str', type: 'text', value: str, min: '1', maxlength: '4' });
	strBoxDiv.appendChild(strBox);

	row.appendChild(keyBoxDiv);
	row.appendChild(strBoxDiv);
	form.insertBefore(row, document.getElementsByClassName('row')[conditions + 2]);
	updateTextboxes();
	conditions++;
}

function updateTextboxes() {
	const textboxes = document.getElementsByClassName('text');
	Array.from(textboxes).forEach((ele) => {
		ele.onclick = function() {
			this.select();
		};
		ele.onkeyup = function(event) {
			if (event.which == 13 || event.keyCode == 13) {
				this.blur();
			}
		};
	});
}

addCondition('3', 'fizz');
addCondition('5', 'buzz');

outputCopy.onclick = function() {
	let range = document.createRange();
	range.selectNode(outputText);
	window.getSelection().removeAllRanges();
	window.getSelection().addRange(range);
	document.execCommand('copy');
};
