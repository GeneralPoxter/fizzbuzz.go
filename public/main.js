const form = document.getElementById('form');
const output = document.getElementById('output');
const outputText = document.getElementById('output-text');
const outputCopy = document.getElementById('output-copy');

form.addEventListener('submit', async function(event) {
	event.preventDefault();
	let params = new URLSearchParams(new FormData(form)).toString();
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

const textboxes = document.getElementsByClassName('text');
Array.from(textboxes).forEach((e) => {
	e.onclick = function() {
		this.select();
	};
});

outputCopy.onclick = function() {
	let range = document.createRange();
	range.selectNode(outputText);
	window.getSelection().removeAllRanges();
	window.getSelection().addRange(range);
	document.execCommand('copy');
};
