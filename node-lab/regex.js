// const rawIgnores = '/\[.*(three).*\n/g'
// const rawIgnores = 'one,two'
const rawIgnores = process.argv[2]
console.log(1, rawIgnores)
let logStr = "[artikow] one\n[artikow] two\n[artikow] three\n[artikow] select 1\n[artikow] select 2\n"
let regexReplace
if (rawIgnores.charAt(0) == '/') {
	const lastSlashIdx = rawIgnores.lastIndexOf('/')
	const regexType = rawIgnores.substring(lastSlashIdx+1)
	const ignores = rawIgnores.substring(0,lastSlashIdx).replace(/\//g,'')
	console.log(ignores,regexType)
	regexReplace = new RegExp(ignores, regexType)
} else {
	const ignores = rawIgnores.replace(/,/g, '|')
	regexReplace = new RegExp(`\\[.*(${ignores}).*\n`, 'g')
}
console.log(3, regexReplace)
logStr = logStr.replace(regexReplace, '')
console.log(logStr)