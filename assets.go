/*
   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
// CONTRIBUTORS AND COPYRIGHT HOLDERS (c) 2013:
// Dag Robøle (dag D0T robole AT gmail D0T com)

package main

// ICO_FavIcon, a favicon in base64 format
var ICO_FavIcon string = `
AAABAAEAEBAAAAEAIABoBAAAFgAAACgAAAAQAAAAIAAAAAEAIAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAABlqclGQm+EjQAAAA4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAcOJDZW
BRcsHAAAAAAAAAAAAAAAAFN/kwZ8xOWNc7vd/UVziL8CDhZxAAAAHAAAAAUAAAAAAAAABQAAACkG
ITuSMI3A/Ao3bCkAAAAAAAAAAG6sxymCyenPe8Pl/3S93v9Fcoe9EUpx2gU1V8EABQlkAAAATAAA
AFMHKVGvL5vG/i2FwOEAAAAAAAAAAAAAAAB1vdzQfsXl/3vD5P9wudv/TH2SyBxJZ8YYd7j/B2Sp
7QgnQqYJQHu8G22x/1rp+f8XQ3GwAAABJQAAABIAAAADc7rZ4H7G5f98xOP/b7jZ/16duf81Z4P/
K4nL/xeI1/8EXKr/EWvJ/0PK5f9gyun/HjhWaAEBARQAAAAMAAAAA3O72uF/x+b/fcXl/2y12P9e
m7b/P2p//yiP0f8wpfT/DXLC/yyc1v9JzO7/M5XO/FmGrhMAAAAAAAAAAAAAAAB4v97hgcjo/37H
5/9xutz/YqPA/0h+l/8imND/JKz0/zGo6v9Zze7/Rb3q/0O66v5Ds+OSNoCeBwAAAAAAAAAAeL7e
4YLJ6P99xef/dLzd/2Cjwv8gbp//HoC9/zqz5v9f0vn/mu/+/5z9/v9z8/7/YOX9/1TT981JrtEk
AAAAAHW83OGCyej/e8Pl/3S83v9Pm8T/Jnmn/0Cv0P9P1Pn/ZtD7/4zs/v+m8P3/mPX9/3Hl9v9V
yej6OZvQzBdeoxtyutrhfcTk/3vD5P9wudv/XrHP/z6euf9at9L/Srnx/27Y/P+c8/7/RJ7V+DSs
2Esvk7EkHmOEBAAAAAAAAAAAc7vZ4X/G5f97w+P/b7jZ/2Wuy/9Ula//X6C//2S65P9y3f3/pfb9
/yd6v7UAAAAAAAAAAAAAAAAAAAAAAAAAAHO72+F/x+b/fcXm/2222P9qrsr/Wpax/2Giwv9qtdj/
dN78/5Tj8/8ffcFiAAAAAAAAAAAAAAAAAAAAAAAAAAB4v97hgcjo/37H5/9rrsz/XZq2/12cuv9l
qcv/abLV/3bU8P9zz+r7InSfEwAAAAAAAAAAAAAAAAAAAAAAAAAAd77e4X7D4f9oqcX/ZKbC/2es
y/9mrdD/Z6/S/2ix1P9zw+H/Ws3rxQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAF+iv+Fgo8D/aK7N
/2qx0f9psdL/aLDS/2au0f9pstT/arXW/2vK5rMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABXnLus
arHQ2nG62r5zvdyhdb/fhXXA32lssctMYZ+0MFCGmBMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAA9/8AAMfnAACBxwAAAAcAAAAPAAAADwAAAAcAAAADAAAAAQAAAB8AAAAfAAAAPwAAAD8AAAA/
AAAAPwAAB/8AAA==`

// PNG_Folder, a base64 encoded folder icon, PNG format
var PNG_Folder string = `
iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAACjElEQVR42nWSTWgTQRiGv/1N00Za
gySkeFAPYkDwoiUHm6RFDzbVBqvVoJBWjyLYg6ARBVsUj0KrBIu4eBBEPChVPBQL4sWDf4h4FEyw
jdUmbLI/szu76zeJlcTGWQZ2Z+d75/3eZzj4Mx7dHN0RCQWmQps2lmVJKFrULSz9rBVLv2qF7yW1
eG76mQZtBrf2olxPZSPhoBLolIGYBHiBB5H3wCdLIEmiJwpchReFgoDiz19/zUxenVdbBGZziezI
UFyRJAkEkYdqTQdVrYGumaAbBGxKwSd4EOzxw+KHUmRyan65RWDmYn82PTygiKIIHK46rguWoYPr
OGATAo5DQTUc8IsuLH5sIzCbi6ODhOLz+QCtYgEWWgQsXceWDADPhYpGoVPy4CVzMP2PwC1sIZ1K
KrJPBo7nsMjEGna6BdQmdcHVqgl+wdEX3i6Hzt94obUI3L6UzI6kEvUM2KILHjjUBooCzAFrZaVS
gw1+YWLX4bvKOgpMID2cVHi2iAQ8fFwMzsL+WQbMzapGzNjROX9bjM0UOFSxbQoUM6CWhU4oOqBQ
JY7ed+ROV1uBmQtI4eCAIiEFwAxYz5Zh1E+2sAX2rdug7xnNd/3HQTybRgpyB1IQGhSYA6Lp2AYG
ihSqxNWv5V8Fnix8WSvzWilgBrLchgIKufhe1iwzNja3mxXipDiXWikMNSiw/0y6QYHUxRiFHxWN
9J+4N44bPuEs4Kz+FchfGcwcOhB/wBYEdpHwJrLg7CYK30pqed/E/f245X2ddHMGZ0/2SaeO7T3e
G+7ZiXc5att2FE/fSm1LYBTYVX7zeSVzOvf4oet63roQ243LZwY7tvR2b9+2uTsaDvqjquE9jY3l
3zXv+Q17WXIurtbLMQAAAABJRU5ErkJggg==`

// PNG_File, a base64 encoded file icon, PNG format
var PNG_File string = `
iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAADWElEQVR42o2SaUxTWRTH/6+VtiJ2
obUUy5NFohFMjNFoUCcZdTIBI1EUw3wwhIjiUhNcgnGJWnEN44gLwpi4oKiJLG6DAdRE64wGQywK
iJBUilArVmpf973M7c34Yb4Y/y959768c373f849DP7ThjVLlKIYfkrYH1BPYVXyGRmpEs7mnOB2
eWIdDlesUCgQjY2N8SORMYtxwHTr/J2/X0fzmOhrUH+tiQTkiSVxjFAkAJ/Px/d0v/Gx+VRVA6t7
Y4hQwMi7W1yCSiHBD6q/dwB7955Pvq17PUQBQ511n58/faX8rSiXBlhGrLCO2jBjZvr/EsPhMDpe
dCE1ncWuHaeyrrW+bKeArifV1ofNz+KLtxQgEAiis+Mt2v/pxLwFs6CYJENCogIMj4cElRzHD9Rg
31ENdm46tv5sk+4SBbQ3/2Hs0velODgnSOMwPTMNHreXJvZ2G+C0uyAaL8TS7IVoa34GbUUptGWn
tYcvPzhEAW3XtT3WL1wmn8/DqOUrluUtRss9HWbPzYDxvQkfBkzIWfEzjIZh8j2Moo35qL3QeHrH
7ze3U0D9uZ36pTkLZwsE4+hJDHnMHy2QysSwcw4wDINEtZKuQ0YzEpOU+LPyRnXpiToNBVw8UtJW
rCn4ldwzzIMmKNUJIFvweAxCwRDcDhd4xF3AHyD/VBR0tqL2aunxuiIKqNhWcLXsYElhOBRGn74b
wvEi+L0+TE5h4XI4ESKNlSnlIEOEiVIxYgQxqKm83ral/Eo2BewvXl5VfnK7Jro3G4fhsNkhIAOV
lJaM92/7kZoxDX2vuiCWSaFOm0IPaLrZ0pGvOTmPArbm/bTn3GXtMW70K3weL7UfDoUgkcvgdXuo
7WHDIJ2FqAOpIh6tf+mMOYXlaRSwLnu+pqb2UFXQ76M1x06Mg5Oz01rjJGLSCx5cdgfGEesCoZCW
cK/hkWllyQmWAtb+Mie3snrPfblCCr2uHUnpKYhXKmAdscDQ0w91KkudREtyEtDkZBZ36h8OrN5c
MZUCcrMyM3aVFfYsWjyX4aw2SIn1bxr9ZIFcNQk+rx9B0swgcWj9YosCGnefaVjDfAssWZ61imVV
eWSYBKTbEaJoMEeSvF6v3+J0eaxuX8BPrtodCEXe3X3e3RvN+xdiqXcg4OWZGwAAAABJRU5ErkJg
gg==`

// TEMPL_Index, HTML template
var TEMPL_Index string = `
<!DOCTYPE HTML>
<html>
<head>
<style>
body { margin: 0; padding: 0; background-color: #c0c0a5; font-size: 15px; }
#header { color:black; text-align:center; padding:5px; }
#nav { width:20%; float:left; padding:5px; }
#section { width:75%; float:left; padding:5px; }
a { text-decoration: none; }
a:link { color: #005555; }
a:visited { color: #005555; }
a:hover { color: #FF0000; }
</style>
</head>
<body>
	<div id="header">	
		<br>
	</div>
	<div id="nav">		
	</div>
	<div id="section">
		{{.}}
	</div>
</body>
</html>`
