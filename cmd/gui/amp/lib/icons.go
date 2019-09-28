package lib

// var ICO = make(map[string]string)
func GetIcons() map[string]string {
	return map[string]string{
		"addressbook": `<svg version="1" xmlns="http://www.w3.org/2000/svg" class="icon" viewBox="0 0 48 48" enable-background="new 0 0 48 48"><path class="light-blue" d="M38,44H12V4h26c2.2,0,4,1.8,4,4v32C42,42.2,40.2,44,38,44z"/><path class="blue" d="M10,4h2v40h-2c-2.2,0-4-1.8-4-4V8C6,5.8,7.8,4,10,4z"/><path class="white" d="M36,24.2c-0.1,4.8-3.1,6.9-5.3,6.7c-0.6-0.1-2.1-0.1-2.9-1.6c-0.8,1-1.8,1.6-3.1,1.6c-2.6,0-3.3-2.5-3.4-3.1 c-0.1-0.7-0.2-1.4-0.1-2.2c0.1-1,1.1-6.5,5.7-6.5c2.2,0,3.5,1.1,3.7,1.3L30,27.2c0,0.3-0.2,1.6,1.1,1.6c2.1,0,2.4-3.9,2.4-4.6 c0.1-1.2,0.3-8.2-7-8.2c-6.9,0-7.9,7.4-8,9.2c-0.5,8.5,6,8.5,7.2,8.5c1.7,0,3.7-0.7,3.9-0.8l0.4,2c-0.3,0.2-2,1.1-4.4,1.1 c-2.2,0-10.1-0.4-9.8-10.8C16.1,23.1,17.4,14,26.6,14C35.8,14,36,22.1,36,24.2z M24.1,25.5c-0.1,1,0,1.8,0.2,2.3 c0.2,0.5,0.6,0.8,1.2,0.8c0.1,0,0.3,0,0.4-0.1c0.2-0.1,0.3-0.1,0.5-0.3c0.2-0.1,0.3-0.3,0.5-0.6c0.2-0.2,0.3-0.6,0.4-1l0.5-5.4 c-0.2-0.1-0.5-0.1-0.7-0.1c-0.5,0-0.9,0.1-1.2,0.3c-0.3,0.2-0.6,0.5-0.9,0.8c-0.2,0.4-0.4,0.8-0.6,1.3S24.2,24.8,24.1,25.5z"/></svg>`,
		"balance":     `<svg version="1" xmlns="http://www.w3.org/2000/svg" class="icon" viewBox="0 0 48 48" enable-background="new 0 0 48 48"><g class="orange"><path d="M38,13c-3.3,0-6-0.9-6-2c0,0.4,0,1.6,0,2c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.4,0-1.6,0-2C44,12.1,41.3,13,38,13 z"/><path d="M38,10c-3.3,0-6-0.9-6-2c0,0.4,0,1.6,0,2c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.4,0-1.6,0-2C44,9.1,41.3,10,38,10z"/><path d="M38,16c-3.3,0-6-0.9-6-2c0,0.4,0,1.6,0,2c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.4,0-1.6,0-2C44,15.1,41.3,16,38,16 z"/><path d="M38,19c-3.3,0-6-0.9-6-2c0,0.4,0,1.6,0,2c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.4,0-1.6,0-2C44,18.1,41.3,19,38,19 z"/><path d="M38,22c-3.3,0-6-0.9-6-2c0,0.4,0,1.6,0,2c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.4,0-1.6,0-2C44,21.1,41.3,22,38,22 z"/><path d="M38,25c-3.3,0-6-0.9-6-2c0,0.4,0,1.6,0,2c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.4,0-1.6,0-2C44,24.1,41.3,25,38,25 z"/><path d="M38,28c-3.3,0-6-0.9-6-2c0,0.4,0,1.6,0,2c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.4,0-1.6,0-2C44,27.1,41.3,28,38,28 z"/><path d="M38,31c-3.3,0-6-0.9-6-2c0,0.4,0,1.6,0,2c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.4,0-1.6,0-2C44,30.1,41.3,31,38,31 z"/><path d="M38,34c-3.3,0-6-0.9-6-2c0,0.4,0,1.6,0,2c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.4,0-1.6,0-2C44,33.1,41.3,34,38,34 z"/><path d="M38,37c-3.3,0-6-0.9-6-2c0,0.4,0,1.6,0,2c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.4,0-1.6,0-2C44,36.1,41.3,37,38,37 z"/><path d="M38,40c-3.3,0-6-0.9-6-2c0,0.4,0,1.6,0,2c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.4,0-1.6,0-2C44,39.1,41.3,40,38,40 z"/></g><g class="yellow"><ellipse cx="38" cy="8" rx="6" ry="2"/><path d="M38,12c-2.8,0-5.1-0.6-5.8-1.5C32.1,10.7,32,10.8,32,11c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.2-0.1-0.3-0.2-0.5 C43.1,11.4,40.8,12,38,12z"/><path d="M38,15c-2.8,0-5.1-0.6-5.8-1.5C32.1,13.7,32,13.8,32,14c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.2-0.1-0.3-0.2-0.5 C43.1,14.4,40.8,15,38,15z"/><path d="M38,18c-2.8,0-5.1-0.6-5.8-1.5C32.1,16.7,32,16.8,32,17c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.2-0.1-0.3-0.2-0.5 C43.1,17.4,40.8,18,38,18z"/><path d="M38,21c-2.8,0-5.1-0.6-5.8-1.5C32.1,19.7,32,19.8,32,20c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.2-0.1-0.3-0.2-0.5 C43.1,20.4,40.8,21,38,21z"/><path d="M38,24c-2.8,0-5.1-0.6-5.8-1.5C32.1,22.7,32,22.8,32,23c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.2-0.1-0.3-0.2-0.5 C43.1,23.4,40.8,24,38,24z"/><path d="M38,27c-2.8,0-5.1-0.6-5.8-1.5C32.1,25.7,32,25.8,32,26c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.2-0.1-0.3-0.2-0.5 C43.1,26.4,40.8,27,38,27z"/><path d="M38,30c-2.8,0-5.1-0.6-5.8-1.5C32.1,28.7,32,28.8,32,29c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.2-0.1-0.3-0.2-0.5 C43.1,29.4,40.8,30,38,30z"/><path d="M38,33c-2.8,0-5.1-0.6-5.8-1.5C32.1,31.7,32,31.8,32,32c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.2-0.1-0.3-0.2-0.5 C43.1,32.4,40.8,33,38,33z"/><path d="M38,36c-2.8,0-5.1-0.6-5.8-1.5C32.1,34.7,32,34.8,32,35c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.2-0.1-0.3-0.2-0.5 C43.1,35.4,40.8,36,38,36z"/><path d="M38,39c-2.8,0-5.1-0.6-5.8-1.5C32.1,37.7,32,37.8,32,38c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.2-0.1-0.3-0.2-0.5 C43.1,38.4,40.8,39,38,39z"/></g><g class="orange"><path d="M10,19c-3.3,0-6-0.9-6-2c0,0.4,0,1.6,0,2c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.4,0-1.6,0-2C16,18.1,13.3,19,10,19 z"/><path d="M10,16c-3.3,0-6-0.9-6-2c0,0.4,0,1.6,0,2c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.4,0-1.6,0-2C16,15.1,13.3,16,10,16 z"/><path d="M10,22c-3.3,0-6-0.9-6-2c0,0.4,0,1.6,0,2c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.4,0-1.6,0-2C16,21.1,13.3,22,10,22 z"/><path d="M10,25c-3.3,0-6-0.9-6-2c0,0.4,0,1.6,0,2c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.4,0-1.6,0-2C16,24.1,13.3,25,10,25 z"/><path d="M10,28c-3.3,0-6-0.9-6-2c0,0.4,0,1.6,0,2c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.4,0-1.6,0-2C16,27.1,13.3,28,10,28 z"/><path d="M10,31c-3.3,0-6-0.9-6-2c0,0.4,0,1.6,0,2c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.4,0-1.6,0-2C16,30.1,13.3,31,10,31 z"/><path d="M10,34c-3.3,0-6-0.9-6-2c0,0.4,0,1.6,0,2c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.4,0-1.6,0-2C16,33.1,13.3,34,10,34 z"/><path d="M10,37c-3.3,0-6-0.9-6-2c0,0.4,0,1.6,0,2c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.4,0-1.6,0-2C16,36.1,13.3,37,10,37 z"/><path d="M10,40c-3.3,0-6-0.9-6-2c0,0.4,0,1.6,0,2c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.4,0-1.6,0-2C16,39.1,13.3,40,10,40 z"/></g><g class="yellow"><ellipse cx="10" cy="14" rx="6" ry="2"/><path d="M10,18c-2.8,0-5.1-0.6-5.8-1.5C4.1,16.7,4,16.8,4,17c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.2-0.1-0.3-0.2-0.5 C15.1,17.4,12.8,18,10,18z"/><path d="M10,21c-2.8,0-5.1-0.6-5.8-1.5C4.1,19.7,4,19.8,4,20c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.2-0.1-0.3-0.2-0.5 C15.1,20.4,12.8,21,10,21z"/><path d="M10,24c-2.8,0-5.1-0.6-5.8-1.5C4.1,22.7,4,22.8,4,23c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.2-0.1-0.3-0.2-0.5 C15.1,23.4,12.8,24,10,24z"/><path d="M10,27c-2.8,0-5.1-0.6-5.8-1.5C4.1,25.7,4,25.8,4,26c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.2-0.1-0.3-0.2-0.5 C15.1,26.4,12.8,27,10,27z"/><path d="M10,30c-2.8,0-5.1-0.6-5.8-1.5C4.1,28.7,4,28.8,4,29c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.2-0.1-0.3-0.2-0.5 C15.1,29.4,12.8,30,10,30z"/><path d="M10,33c-2.8,0-5.1-0.6-5.8-1.5C4.1,31.7,4,31.8,4,32c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.2-0.1-0.3-0.2-0.5 C15.1,32.4,12.8,33,10,33z"/><path d="M10,36c-2.8,0-5.1-0.6-5.8-1.5C4.1,34.7,4,34.8,4,35c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.2-0.1-0.3-0.2-0.5 C15.1,35.4,12.8,36,10,36z"/><path d="M10,39c-2.8,0-5.1-0.6-5.8-1.5C4.1,37.7,4,37.8,4,38c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.2-0.1-0.3-0.2-0.5 C15.1,38.4,12.8,39,10,39z"/></g><g class="orange"><path d="M24,28c-3.3,0-6-0.9-6-2c0,0.4,0,1.6,0,2c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.4,0-1.6,0-2C30,27.1,27.3,28,24,28 z"/><path d="M24,25c-3.3,0-6-0.9-6-2c0,0.4,0,1.6,0,2c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.4,0-1.6,0-2C30,24.1,27.3,25,24,25 z"/><path d="M24,31c-3.3,0-6-0.9-6-2c0,0.4,0,1.6,0,2c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.4,0-1.6,0-2C30,30.1,27.3,31,24,31 z"/><path d="M24,34c-3.3,0-6-0.9-6-2c0,0.4,0,1.6,0,2c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.4,0-1.6,0-2C30,33.1,27.3,34,24,34 z"/><path d="M24,37c-3.3,0-6-0.9-6-2c0,0.4,0,1.6,0,2c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.4,0-1.6,0-2C30,36.1,27.3,37,24,37 z"/><path d="M24,40c-3.3,0-6-0.9-6-2c0,0.4,0,1.6,0,2c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.4,0-1.6,0-2C30,39.1,27.3,40,24,40 z"/></g><g class="yellow"><ellipse cx="24" cy="23" rx="6" ry="2"/><path d="M24,27c-2.8,0-5.1-0.6-5.8-1.5C18.1,25.7,18,25.8,18,26c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.2-0.1-0.3-0.2-0.5 C29.1,26.4,26.8,27,24,27z"/><path d="M24,30c-2.8,0-5.1-0.6-5.8-1.5C18.1,28.7,18,28.8,18,29c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.2-0.1-0.3-0.2-0.5 C29.1,29.4,26.8,30,24,30z"/><path d="M24,33c-2.8,0-5.1-0.6-5.8-1.5C18.1,31.7,18,31.8,18,32c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.2-0.1-0.3-0.2-0.5 C29.1,32.4,26.8,33,24,33z"/><path d="M24,36c-2.8,0-5.1-0.6-5.8-1.5C18.1,34.7,18,34.8,18,35c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.2-0.1-0.3-0.2-0.5 C29.1,35.4,26.8,36,24,36z"/><path d="M24,39c-2.8,0-5.1-0.6-5.8-1.5C18.1,37.7,18,37.8,18,38c0,1.1,2.7,2,6,2s6-0.9,6-2c0-0.2-0.1-0.3-0.2-0.5 C29.1,38.4,26.8,39,24,39z"/></g></svg>`,
		"blocks":      `<svg version="1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" enable-background="new 0 0 48 48"><g class="purple" ><path d="M38,7H10C8.9,7,8,7.9,8,9v6c0,1.1,0.9,2,2,2h28c1.1,0,2-0.9,2-2V9C40,7.9,39.1,7,38,7z"/><path d="M38,19H10c-1.1,0-2,0.9-2,2v6c0,1.1,0.9,2,2,2h28c1.1,0,2-0.9,2-2v-6C40,19.9,39.1,19,38,19z"/><path d="M38,31H10c-1.1,0-2,0.9-2,2v6c0,1.1,0.9,2,2,2h28c1.1,0,2-0.9,2-2v-6C40,31.9,39.1,31,38,31z"/></g></svg>`,
		"help":        `<svg version="1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" enable-background="new 0 0 48 48"><path  class="orange"  d="M32,15v28H10c-2.2,0-4-1.8-4-4V15H32z"/><path  class="red" d="M14,5v34c0,2.2-1.8,4-4,4h29c2.2,0,4-1.8,4-4V5H14z"/><g  class="orange"><rect x="20" y="10" width="18" height="4"/><rect x="20" y="17" width="8" height="2"/><rect x="30" y="17" width="8" height="2"/><rect x="20" y="21" width="8" height="2"/><rect x="30" y="21" width="8" height="2"/><rect x="20" y="25" width="8" height="2"/><rect x="30" y="25" width="8" height="2"/><rect x="20" y="29" width="8" height="2"/><rect x="30" y="29" width="8" height="2"/><rect x="20" y="33" width="8" height="2"/><rect x="30" y="33" width="8" height="2"/><rect x="20" y="37" width="8" height="2"/><rect x="30" y="37" width="8" height="2"/></g></svg>`,
		"history":     `<svg version="1" xmlns="http://www.w3.org/2000/svg" class="icon" viewBox="0 0 48 48" enable-background="new 0 0 48 48"><path class="white" d="M5,38V14h38v24c0,2.2-1.8,4-4,4H9C6.8,42,5,40.2,5,38z"/><path class="red" d="M43,10v6H5v-6c0-2.2,1.8-4,4-4h30C41.2,6,43,7.8,43,10z"/><g class="white"><circle cx="33" cy="10" r="3"/><circle cx="15" cy="10" r="3"/></g><g class="gray"><path d="M33,3c-1.1,0-2,0.9-2,2v5c0,1.1,0.9,2,2,2s2-0.9,2-2V5C35,3.9,34.1,3,33,3z"/><path d="M15,3c-1.1,0-2,0.9-2,2v5c0,1.1,0.9,2,2,2s2-0.9,2-2V5C17,3.9,16.1,3,15,3z"/></g><g class="gray"><rect x="13" y="20" width="4" height="4"/><rect x="19" y="20" width="4" height="4"/><rect x="25" y="20" width="4" height="4"/><rect x="31" y="20" width="4" height="4"/><rect x="13" y="26" width="4" height="4"/><rect x="19" y="26" width="4" height="4"/><rect x="25" y="26" width="4" height="4"/><rect x="31" y="26" width="4" height="4"/><rect x="13" y="32" width="4" height="4"/><rect x="19" y="32" width="4" height="4"/><rect x="25" y="32" width="4" height="4"/><rect x="31" y="32" width="4" height="4"/></g></svg>`,
		"info":        `<svg version="1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" enable-background="new 0 0 48 48"><circle class="blue" cx="24" cy="24" r="21"/><rect x="22" y="22" class="white"  width="4" height="11"/><circle class="white"  cx="24" cy="16.5" r="2.5"/></svg>`,
		"link":        `<svg version="1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" enable-background="new 0 0 48 48"><g class="blue"><path d="M38,13h-3c-5.5,0-10,4.5-10,10s4.5,10,10,10h3c5.5,0,10-4.5,10-10S43.5,13,38,13z M38,29h-3 c-3.3,0-6-2.7-6-6s2.7-6,6-6h3c3.3,0,6,2.7,6,6S41.3,29,38,29z"/><path d="M13,13h-3C4.5,13,0,17.5,0,23s4.5,10,10,10h3c5.5,0,10-4.5,10-10S18.5,13,13,13z M13,29h-3 c-3.3,0-6-2.7-6-6s2.7-6,6-6h3c3.3,0,6,2.7,6,6S16.3,29,13,29z"/></g><path class="light-blue" d="M33,21H15c-1.1,0-2,0.9-2,2s0.9,2,2,2h18c1.1,0,2-0.9,2-2S34.1,21,33,21z"/></svg>`,
		"linkOff":     `<svg version="1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" enable-background="new 0 0 48 48"><g class="blue" ><path d="M17.5,27c-1.1,1.2-2.7,2-4.5,2h-3c-3.3,0-6-2.7-6-6s2.7-6,6-6h3c1.8,0,3.4,0.8,4.5,2h4.7 c-1.5-3.5-5.1-6-9.2-6h-3C4.5,13,0,17.5,0,23s4.5,10,10,10h3c4.1,0,7.6-2.5,9.2-6H17.5z"/><path d="M38,13h-3c-4.1,0-7.6,2.5-9.2,6h4.7c1.1-1.2,2.7-2,4.5-2h3c3.3,0,6,2.7,6,6s-2.7,6-6,6h-3 c-1.8,0-3.4-0.8-4.5-2h-4.7c1.5,3.5,5.1,6,9.2,6h3c5.5,0,10-4.5,10-10S43.5,13,38,13z"/></g><g class="light-blue"><polygon points="19.5,4 16,6 22.1,14.1 23.4,13.3"/><polygon points="28.5,4 32,6 25.9,14.1 24.6,13.3"/><polygon points="28.5,44 32,42 25.9,33.9 24.6,34.7"/><polygon points="19.5,44 16,42 22.1,33.9 23.4,34.7"/></g></svg>`,
		"loading":     `<svg version="1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" enable-background="new 0 0 48 48"><g class="green anim"><polygon points="31,8 42.9,9.6 33.1,19.4"/><polygon points="17,40 5.1,38.4 14.9,28.6"/><polygon points="8,17 9.6,5.1 19.4,14.9"/><path d="M9.3,21.2L5.1,22C5,22.7,5,23.3,5,24c0,4.6,1.6,9,4.6,12.4l3-2.6C10.3,31.1,9,27.6,9,24 C9,23.1,9.1,22.1,9.3,21.2z"/><path d="M24,5c-5.4,0-10.2,2.3-13.7,5.9l2.8,2.8C15.9,10.8,19.7,9,24,9c0.9,0,1.9,0.1,2.8,0.3l0.7-3.9 C26.4,5.1,25.2,5,24,5z"/><path d="M38.7,26.8l4.2-0.8c0.1-0.7,0.1-1.3,0.1-2c0-4.4-1.5-8.7-4.3-12.1l-3.1,2.5c2.2,2.7,3.4,6.1,3.4,9.5 C39,24.9,38.9,25.9,38.7,26.8z"/><path d="M34.9,34.3C32.1,37.2,28.3,39,24,39c-0.9,0-1.9-0.1-2.8-0.3l-0.7,3.9c1.2,0.2,2.4,0.3,3.5,0.3 c5.4,0,10.2-2.3,13.7-5.9L34.9,34.3z"/><polygon points="40,31 38.4,42.9 28.6,33.1"/></g></svg>`,
		"logo":        `<svg xmlns="http://www.w3.org/2000/svg" id="parallel" viewBox="0 0 108 128" width="108" height="128"><path class="logofill" d="M77.08,2.55c3.87,1.03 6.96,2.58 10.32,4.64c5.93,3.87 10.58,8.51 14.19,14.71c3.87,6.19 5.42,13.16 5.42,20.64c0,7.22 -1.81,14.18 -5.41,20.37c-3.61,6.45 -8.25,11.35 -14.19,14.96c-3.35,2.06 -6.96,3.87 -10.32,4.9c-3.87,1.03 -7.74,1.55 -11.61,1.55v-14.45c6.96,-0.26 13.42,-2.58 19.09,-8c5.67,-5.42 8.51,-11.87 8.51,-19.61c0,-7.74 -2.58,-14.19 -7.99,-19.6c-5.42,-5.42 -11.86,-8 -19.6,-8c-7.74,0 -14.44,2.58 -19.6,8c-5.42,5.42 -8,11.87 -8,19.6l0,85.9c-3.1,-3.1 -7.99,-7.74 -13.93,-13.67v-72.23c0,-3.87 0.52,-7.73 1.55,-11.35c1.03,-3.87 2.58,-7.22 4.64,-10.32c3.87,-5.93 8.52,-10.58 14.71,-14.45c6.19,-3.61 13.16,-5.16 20.64,-5.16c3.87,0 8,0.52 11.61,1.55zM78.37,42.28c0,7.22 -5.93,13.16 -13.15,13.16c-7.48,0.26 -13.16,-5.68 -13.16,-13.16c0,-7.22 5.94,-13.16 13.16,-13.16c7.22,0 13.15,5.93 13.15,13.16zM13.63,37.12l0,69.39c-6.19,-6.19 -11.09,-10.83 -13.93,-13.93l0,-55.46z" /></svg>`,
		"overview":    `<svg version="1" xmlns="http://www.w3.org/2000/svg" class="icon"  viewBox="0 0 48 48" enable-background="new 0 0 48 48"><polygon class="white" points="42,39 6,39 6,23 24,6 42,23"/><g class="blue"><polygon points="39,21 34,16 34,9 39,9"/><rect x="6" y="39" width="36" height="5"/></g><polygon class="red" points="24,4.3 4,22.9 6,25.1 24,8.4 42,25.1 44,22.9"/><rect x="18" y="28" class="red" width="12" height="16"/><rect x="21" y="17" class="blue" width="6" height="6"/><path class="green" d="M27.5,35.5c-0.3,0-0.5,0.2-0.5,0.5v2c0,0.3,0.2,0.5,0.5,0.5S28,38.3,28,38v-2C28,35.7,27.8,35.5,27.5,35.5z"/></svg>`,
		"peers":       `<svg version="1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" enable-background="new 0 0 48 48"><polygon class="gray" points="39.4,23 38.6,19 26,21.6 26,8 22,8 22,20.3 8.1,11.3 5.9,14.7 21.1,24.5 9.4,39.8 12.6,42.2 23.9,27.4 32.3,40.1 35.7,37.9 27.3,25.4"/><circle class="blue" cx="24" cy="24" r="7"/><g class="light-blue"><circle cx="24" cy="8" r="5"/><circle cx="39" cy="21" r="5"/><circle cx="7" cy="13" r="5"/><circle cx="11" cy="41" r="5"/><circle cx="34" cy="39" r="5"/></g></svg>`,
		"receive":     `<svg version="1" xmlns="http://www.w3.org/2000/svg" class="icon" viewBox="0 0 48 48" enable-background="new 0 0 48 48"><g class="blue"><polygon points="24,37.1 13,24 35,24"/><rect x="20" y="4" width="8" height="4"/><rect x="20" y="10" width="8" height="4"/><rect x="20" y="16" width="8" height="11"/><rect x="6" y="40" width="36" height="4"/></g></svg>`,
		"received":    `<svg version="1" xmlns="http://www.w3.org/2000/svg" class="icon" viewBox="0 0 48 48" enable-background="new 0 0 48 48"><path class="gray" d="M7,40V8c0-2.2,1.8-4,4-4h24c2.2,0,4,1.8,4,4v32c0,2.2-1.8,4-4,4H11C8.8,44,7,42.2,7,40z"/><g class="green"><polygon points="13.3,24 24,15 24,33"/><rect x="19" y="21" width="23" height="6"/></g></svg>`,
		"send":        `<svg version="1" xmlns="http://www.w3.org/2000/svg" class="icon"  viewBox="0 0 48 48" enable-background="new 0 0 48 48"><g class="green"><polygon points="24,10.9 35,24 13,24"/><rect x="20" y="40" width="8" height="4"/><rect x="20" y="34" width="8" height="4"/><rect x="20" y="21" width="8" height="11"/><rect x="6" y="4" width="36" height="4"/></g></svg>`,
		"sent":        `<svg version="1" xmlns="http://www.w3.org/2000/svg" class="icon" viewBox="0 0 48 48" enable-background="new 0 0 48 48"><path class="gray" d="M7,40V8c0-2.2,1.8-4,4-4h24c2.2,0,4,1.8,4,4v32c0,2.2-1.8,4-4,4H11C8.8,44,7,42.2,7,40z"/><g class="red"><polygon points="42.7,24 32,33 32,15"/><rect x="14" y="21" width="23" height="6"/></g></svg>`,
		"settings":    `<svg version="1" xmlns="http://www.w3.org/2000/svg" class="icon"  viewBox="0 0 48 48" enable-background="new 0 0 48 48"><path class="gray" d="M39.6,27.2c0.1-0.7,0.2-1.4,0.2-2.2s-0.1-1.5-0.2-2.2l4.5-3.2c0.4-0.3,0.6-0.9,0.3-1.4L40,10.8 c-0.3-0.5-0.8-0.7-1.3-0.4l-5,2.3c-1.2-0.9-2.4-1.6-3.8-2.2l-0.5-5.5c-0.1-0.5-0.5-0.9-1-0.9h-8.6c-0.5,0-1,0.4-1,0.9l-0.5,5.5 c-1.4,0.6-2.7,1.3-3.8,2.2l-5-2.3c-0.5-0.2-1.1,0-1.3,0.4l-4.3,7.4c-0.3,0.5-0.1,1.1,0.3,1.4l4.5,3.2c-0.1,0.7-0.2,1.4-0.2,2.2 s0.1,1.5,0.2,2.2L4,30.4c-0.4,0.3-0.6,0.9-0.3,1.4L8,39.2c0.3,0.5,0.8,0.7,1.3,0.4l5-2.3c1.2,0.9,2.4,1.6,3.8,2.2l0.5,5.5 c0.1,0.5,0.5,0.9,1,0.9h8.6c0.5,0,1-0.4,1-0.9l0.5-5.5c1.4-0.6,2.7-1.3,3.8-2.2l5,2.3c0.5,0.2,1.1,0,1.3-0.4l4.3-7.4 c0.3-0.5,0.1-1.1-0.3-1.4L39.6,27.2z M24,35c-5.5,0-10-4.5-10-10c0-5.5,4.5-10,10-10c5.5,0,10,4.5,10,10C34,30.5,29.5,35,24,35z"/><path class="black" d="M24,13c-6.6,0-12,5.4-12,12c0,6.6,5.4,12,12,12s12-5.4,12-12C36,18.4,30.6,13,24,13z M24,30 c-2.8,0-5-2.2-5-5c0-2.8,2.2-5,5-5s5,2.2,5,5C29,27.8,26.8,30,24,30z"/></svg>`,
		"txnumber":    `<svg version="1" xmlns="http://www.w3.org/2000/svg" class="icon" viewBox="0 0 48 48" enable-background="new 0 0 48 48"><g class="blue"><polygon points="17.8,18.1 10.4,25.4 6.2,21.3 4,23.5 10.4,29.9 20,20.3"/><polygon points="17.8,5.1 10.4,12.4 6.2,8.3 4,10.5 10.4,16.9 20,7.3"/><polygon points="17.8,31.1 10.4,38.4 6.2,34.3 4,36.5 10.4,42.9 20,33.3"/></g><g class="light-blue"><rect x="24" y="22" width="20" height="4"/><rect x="24" y="9" width="20" height="4"/><rect x="24" y="35" width="20" height="4"/></g></svg>`,
		"unconfirmed": `<svg version="1" xmlns="http://www.w3.org/2000/svg" class="icon" viewBox="0 0 48 48" enable-background="new 0 0 48 48"><circle class="blue" cx="17" cy="17" r="14"/><circle class="white" cx="17" cy="17" r="11"/><rect x="16" y="8" width="2" height="9"/><rect x="18.2" y="16" transform="matrix(-.707 .707 -.707 -.707 46.834 19.399)" width="2.4" height="6.8"/><circle cx="17" cy="17" r="2"/><circle class="blue" cx="17" cy="17" r="1"/><path  class="yellow" d="M11.9,42l14.4-24.1c0.8-1.3,2.7-1.3,3.4,0L44.1,42c0.8,1.3-0.2,3-1.7,3H13.6C12.1,45,11.1,43.3,11.9,42z"/><path class="black" d="M26.4,39.9c0-0.2,0-0.4,0.1-0.6s0.2-0.3,0.3-0.5s0.3-0.2,0.5-0.3s0.4-0.1,0.6-0.1s0.5,0,0.7,0.1 s0.4,0.2,0.5,0.3s0.2,0.3,0.3,0.5s0.1,0.4,0.1,0.6s0,0.4-0.1,0.6s-0.2,0.3-0.3,0.5s-0.3,0.2-0.5,0.3s-0.4,0.1-0.7,0.1 s-0.5,0-0.6-0.1s-0.4-0.2-0.5-0.3s-0.2-0.3-0.3-0.5S26.4,40.1,26.4,39.9z M29.2,36.8h-2.3L26.5,27h3L29.2,36.8z"/></svg>`,
		"light":       `<svg version="1" xmlns="http://www.w3.org/2000/svg" class="icon" viewBox="0 0 48 48" enable-background="new 0 0 48 48"><circle class="yellow" cx="24" cy="22" r="20"/><path class="orange" d="M37,22c0-7.7-6.6-13.8-14.5-12.9c-6,0.7-10.8,5.5-11.4,11.5c-0.5,4.6,1.4,8.7,4.6,11.3 c1.4,1.2,2.3,2.9,2.3,4.8V37h12v-0.1c0-1.8,0.8-3.6,2.2-4.8C35.1,29.7,37,26.1,37,22z"/><path class="yellow" d="M30.6,20.2l-3-2c-0.3-0.2-0.8-0.2-1.1,0L24,19.8l-2.4-1.6c-0.3-0.2-0.8-0.2-1.1,0l-3,2 c-0.2,0.2-0.4,0.4-0.4,0.7s0,0.6,0.2,0.8l3.8,4.7V37h2V26c0-0.2-0.1-0.4-0.2-0.6l-3.3-4.1l1.5-1l2.4,1.6c0.3,0.2,0.8,0.2,1.1,0 l2.4-1.6l1.5,1l-3.3,4.1C25.1,25.6,25,25.8,25,26v11h2V26.4l3.8-4.7c0.2-0.2,0.3-0.5,0.2-0.8S30.8,20.3,30.6,20.2z"/><circle class="blue" cx="24" cy="44" r="3"/><path class="white" d="M26,45h-4c-2.2,0-4-1.8-4-4v-5h12v5C30,43.2,28.2,45,26,45z"/><g class="gray"><path d="M30,41l-11.6,1.6c0.3,0.7,0.9,1.4,1.6,1.8l9.4-1.3C29.8,42.5,30,41.8,30,41z"/><polygon points="18,38.7 18,40.7 30,39 30,37"/></g></svg>`,
		"nolight":     `<svg version="1" xmlns="http://www.w3.org/2000/svg" class="icon" viewBox="0 0 48 48" enable-background="new 0 0 48 48"><path class="yellow" d="M37,22c0-7.7-6.6-13.8-14.5-12.9c-6,0.7-10.8,5.5-11.4,11.5c-0.5,4.6,1.4,8.7,4.6,11.3 c1.4,1.2,2.3,2.9,2.3,4.8V37h12v-0.1c0-1.8,0.8-3.6,2.2-4.8C35.1,29.7,37,26.1,37,22z"/><path class="white" d="M30.6,20.2l-3-2c-0.3-0.2-0.8-0.2-1.1,0L24,19.8l-2.4-1.6c-0.3-0.2-0.8-0.2-1.1,0l-3,2 c-0.2,0.2-0.4,0.4-0.4,0.7s0,0.6,0.2,0.8l3.8,4.7V37h2V26c0-0.2-0.1-0.4-0.2-0.6l-3.3-4.1l1.5-1l2.4,1.6c0.3,0.2,0.8,0.2,1.1,0 l2.4-1.6l1.5,1l-3.3,4.1C25.1,25.6,25,25.8,25,26v11h2V26.4l3.8-4.7c0.2-0.2,0.3-0.5,0.2-0.8S30.8,20.3,30.6,20.2z"/><circle class="blue" cx="24" cy="44" r="3"/><path class="blue" d="M26,45h-4c-2.2,0-4-1.8-4-4v-5h12v5C30,43.2,28.2,45,26,45z"/><g class="white"><path d="M30,41l-11.6,1.6c0.3,0.7,0.9,1.4,1.6,1.8l9.4-1.3C29.8,42.5,30,41.8,30,41z"/><polygon points="18,38.7 18,40.7 30,39 30,37"/></g><rect x="22" y="-2.9" transform="matrix(.707 -.707 .707 .707 -9.941 24)" class="gray"  width="4" height="53.7"/></svg>`,
		"search":      `<svg version="1" xmlns="http://www.w3.org/2000/svg" class="icon" viewBox="0 0 48 48" enable-background="new 0 0 48 48"><g fill="#616161"><rect x="34.6" y="28.1" transform="matrix(.707 -.707 .707 .707 -15.154 36.586)" width="4" height="17"/><circle cx="20" cy="20" r="16"/></g><rect x="36.2" y="32.1" transform="matrix(.707 -.707 .707 .707 -15.839 38.239)" fill="#37474F" width="4" height="12.3"/><circle fill="#64B5F6" cx="20" cy="20" r="13"/><path fill="#BBDEFB" d="M26.9,14.2c-1.7-2-4.2-3.2-6.9-3.2s-5.2,1.2-6.9,3.2c-0.4,0.4-0.3,1.1,0.1,1.4c0.4,0.4,1.1,0.3,1.4-0.1 C16,13.9,17.9,13,20,13s4,0.9,5.4,2.5c0.2,0.2,0.5,0.4,0.8,0.4c0.2,0,0.5-0.1,0.6-0.2C27.2,15.3,27.2,14.6,26.9,14.2z"/></svg>`,
		"copy":        `<svg version="1" xmlns="http://www.w3.org/2000/svg" class="icon" viewBox="0 0 48 48" enable-background="new 0 0 48 48"><polygon class="light-blue" points="40,45 8,45 8,3 30,3 40,13"/><polygon class="blue" points="38.5,14 29,14 29,4.5"/><g class="white"><rect x="16" y="21" width="17" height="2"/><rect x="16" y="25" width="13" height="2"/><rect x="16" y="29" width="17" height="2"/><rect x="16" y="33" width="13" height="2"/></g></svg>`,
		"ok":          `<svg version="1" xmlns="http://www.w3.org/2000/svg" class="icon"  viewBox="0 0 48 48" enable-background="new 0 0 48 48"><circle fill="#4CAF50" cx="24" cy="24" r="21"/><polygon fill="#CCFF90" points="34.6,14.6 21,28.2 15.4,22.6 12.6,25.4 21,33.8 37.4,17.4"/></svg>`,
	}
}
