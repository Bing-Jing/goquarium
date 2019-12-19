package goquarium

import (
	"github.com/ansoni/termination"
)

/**
 * This is all of our shapes!
 */

var surfaceShape = termination.Shape{
	"default": []string{
		"~~~~",
	},
}

var waterMask = termination.Shape{
	"default": []string{
		"bbbb",
	},
	"a": []string{
		"cccc\n" +
			"cccc\n" +
			"cccc",
	},
	"b": []string{
		"cccc\n" +
			"cccc\n" +
			"cccc",
	},
	"c": []string{
		"cccc\n" +
			"cccc\n" +
			"cccc",
	},
	"d": []string{
		"cccc\n" +
			"cccc\n" +
			"cccc",
	},
}

var rippleShape = termination.Shape{
	"a": []string{
		"^^^^\n" +
			"^^^^\n" +
			"^^  ",
	},
	"b": []string{
		" ^^^\n" +
			"    \n" +
			"    ",
	},
	"c": []string{
		"  ^^\n" +
			"^^  \n" +
			"   ^",
	},
	"d": []string{
		"    \n" +
			"    \n" +
			"    ",
	},
	"e": []string{
		" ^  \n" +
			"  ^ \n" +
			"   ^",
	},
}

var castleShape = termination.Shape{
	"default": []string{
		`               T~~
               |
              /^\
             /   \
 _   _   _  /     \  _   _   _
[ ]_[ ]_[ ]/ _   _ \[ ]_[ ]_[ ]
|_=__-_ =_|_[ ]_[ ]_|_=-___-__|
 | _- =  | =_ = _    |= _=   |
 |= -[]  |- = _ =    |_-=_[] |
 | =_    |= - ___    | =_ =  |
 |=  []- |-  /| |\   |=_ =[] |
 |- =_   | =| | | |  |- = -  |
 |_______|__|_|_|_|__|_______|`,
		`               T~
               |
              /^\
             /   \
 _   _   _  /     \  _   _   _
[ ]_[ ]_[ ]/ _   _ \[ ]_[ ]_[ ]
|_=__-_ =_|_[ ]_[ ]_|_=-___-__|
 | _- =  | =_ = _    |= _=   |
 |= -[]  |- = _ =    |_-=_[] |
 | =_    |= - ___    | =_ =  |
 |=  []- |-  /| |\   |=_ =[] |
 |- =_   | =| | | |  |- = -  |
 |_______|__|_|_|_|__|_______|`,
	},
}

var castleMask = termination.Shape{
	"default": []string{
		`               Trr
               |
              yyy
             y   y
 _   _   _  y     y  _   _   _
[ ]_[ ]_[ ]y _   _ y[ ]_[ ]_[ ]
|_=__-_ =_|_[ ]_[ ]_|_=-___-__|
 | _- =  | =_ = _    |= _=   |
 |= -[]  |- = _ =    |_-=_[] |
 | =_    |= - yyy    | =_ =  |
 |=  []- |-  yy yy   |=_ =[] |
 |- =_   | =y y y y  |- = -  |
 |_______|__y_y_y_y__|_______|`,
		`               Tr
               |
              yyy
             y   y
 _   _   _  y     y  _   _   _
[ ]_[ ]_[ ]y _   _ y[ ]_[ ]_[ ]
|_=__-_ =_|_[ ]_[ ]_|_=-___-__|
 | _- =  | =_ = _    |= _=   |
 |= -[]  |- = _ =    |_-=_[] |
 | =_    |= - yyy    | =_ =  |
 |=  []- |-  yy yy   |=_ =[] |
 |- =_   | =y y y y  |- = -  |
 |_______|__y_y_y_y__|_______|`,
	},
}

var seaweedShape = termination.Shape{
	"a": []string{
		"( ",
		" )",
	},
	"b": []string{
		" )",
		"( ",
	},
}

var seaweedMask = termination.Shape{
	"a": []string{
		"g ",
		" g",
	},
	"b": []string{
		" g",
		"g ",
	},
}

var bubbleShape = termination.Shape{
	"default": []string{
		".",
		".",
		".",
		".",
		"o",
		"o",
		"o",
		"o",
		"O",
		"O",
		"O",
	},
}

var fishShapes = []termination.Shape{
	fish1Shape,
	fish2Shape,
	fish3Shape,
	fish4Shape,
}

var fishMasks = []termination.Shape{
	fish1Mask,
	fish2Mask,
	fish3Mask,
	fish4Mask,
}

var fish1Shape = termination.Shape{
	"left": []string{
		`???/
??/ \
<')_=<
??\_/
???\`,
	},
	"right": []string{
		`?\
?/ \
>=_('>
?\_/
??/`,
	},
}

var fish1Mask = termination.Shape{
	"left": []string{
		`   m
  r r
y'brgg
  rrr
   y`,
	},
	"right": []string{
		` m
 r r
ggrb'y
rrrr
  y`,
	},
}

var fish2Shape = termination.Shape{
	"left": []string{
		`?__
/o \/
\__/\`,
	},
	"right": []string{
		`??__
\/ o\
/\__/`,
	},
}

var fish2Mask = termination.Shape{
	"left": []string{
		`?gg
gW gb
ggggb`,
	},
	"right": []string{
		`??bb
bb Wb
bbbbb`,
	},
}

var fish3Shape = termination.Shape{
	"left": []string{
		" /,\n" +
			"<')=<\n" +
			" \\`",
	},
	"right": []string{
		" ,\\\n" +
			">=('>\n" +
			" '/",
	},
}

var fish3Mask = termination.Shape{
	"left": []string{
		` rg
b'ycc
 cg`,
	},
	"right": []string{
		` gr
ccy'b
 gc`,
	},
}
var fish4Shape = termination.Shape{
	"left": []string{
		"????_,--,\n" +
			"?.-'---./_    __\n" +
			"/o \\\\     \"-.' /\n" +
			"\\  //    _.-'._\\\n" +
			"?`\"\\)--\"`",
	},
	"right": []string{
		"???????,--,_\n" +
			"__????_\\.---'-.\n" +
			"\\ '.-\"     // o\\\n" +
			"/_.'-._    \\\\  /\n" +
			"???????`\"--(/\"`",
	},
}

var fish4Mask = termination.Shape{
	"left": []string{
		"????ggggg\n" +
			"?rrgrrrrgr    gg\n" +
			"ro rr     rrgg g\n" +
			"r  rr    rrrrggg\n" +
			"?rrrggrrrr",
	},
	"right": []string{
		"???????ggggg\n" +
			"gg????ygyyyygyy\n" +
			"g gyyyy    yy oy\n" +
			"gggyyyy    yy  y\n" +
			"???????yyyyggyy",
	},
}

var whaleShape = termination.Shape{
	"left": []string{
		"                \n" +
			"                \n" +
			"                \n" +
			"    :-----.\n" +
			"  .'       `.\n" +
			" / (o)       \\????,\n" +
			"(__,          \\_.'/\n",
		"                \n" +
			"                \n" +
			"                \n" +
			"    :-----.\n" +
			"  .'       `.\n" +
			" / (o)       \\????,\n" +
			"(__,          \\_.'/\n",
		"                \n" +
			"                \n" +
			"    :\n" +
			"    :-----.\n" +
			"  .'       `.\n" +
			" / (o)       \\????,\n" +
			"(__,          \\_.'/\n",
		"                \n" +
			"    :\n" +
			"    :\n" +
			"    :-----.\n" +
			"  .'       `.\n" +
			" / (o)       \\????,\n" +
			"(__,          \\_.'/\n",
		"   . .\n" +
			"    :\n" +
			"    :\n" +
			"    :-----.\n" +
			"  .'       `.\n" +
			" / (o)       \\????,\n" +
			"(__,          \\_.'/\n",
		"   . .\n" +
			" '.-:-.`\n" +
			"    :\n" +
			"    :-----.\n" +
			"  .'       `.\n" +
			" / (o)       \\????,\n" +
			"(__,          \\_.'/\n",
		"   . .\n" +
			" '.-:-.`\n" +
			" '  :  '\n" +
			"    :-----.\n" +
			"  .'       `.\n" +
			" / (o)       \\????,\n" +
			"(__,          \\_.'/\n",
	},
	"right": []string{
		"                \n" +
			"                \n" +
			"                \n" +
			"        .-----:\n" +
			"      .'       `.\n" +
			",????/       (o) \\\n" +
			"\\`._/          ,__)\n",
		"                \n" +
			"                \n" +
			"                \n" +
			"        .-----:\n" +
			"      .'       `.\n" +
			",????/       (o) \\\n" +
			"\\`._/          ,__)\n",
		"                \n" +
			"                \n" +
			"              :\n" +
			"        .-----:\n" +
			"      .'       `.\n" +
			",????/       (o) \\\n" +
			"\\`._/          ,__)\n",
		"                \n" +
			"              :\n" +
			"              :\n" +
			"        .-----:\n" +
			"      .'       `.\n" +
			",????/       (o) \\\n" +
			"\\`._/          ,__)\n",
		"             . .\n" +
			"              :\n" +
			"              :\n" +
			"        .-----:\n" +
			"      .'       `.\n" +
			",????/       (o) \\\n" +
			"\\`._/          ,__)\n",
		"             . .\n" +
			"           '.-:-.`\n" +
			"              :\n" +
			"        .-----:\n" +
			"      .'       `.\n" +
			",????/       (o) \\\n" +
			"\\`._/          ,__)\n",
		"             . .\n" +
			"           '.-:-.`\n" +
			"           '  :  '\n" +
			"        .-----:\n" +
			"      .'       `.\n" +
			",????/       (o) \\\n" +
			"\\`._/          ,__)\n",
	},
}

var whaleMask = termination.Shape{
	"left": []string{
		"                \n" +
			"                \n" +
			"                \n" +
			"    :-----.\n" +
			"  .'       `.\n" +
			" / (W)       \\????,\n" +
			"(__,          \\_.'/\n",
		"                \n" +
			"                \n" +
			"                \n" +
			"    :-----.\n" +
			"  .'       `.\n" +
			" / (W)       \\????,\n" +
			"(__,          \\_.'/\n",
		"                \n" +
			"                \n" +
			"    c\n" +
			"    :-----.\n" +
			"  .'       `.\n" +
			" / (W)       \\????,\n" +
			"(__,          \\_.'/\n",
		"                \n" +
			"    c\n" +
			"    c\n" +
			"    :-----.\n" +
			"  .'       `.\n" +
			" / (W)       \\????,\n" +
			"(__,          \\_.'/\n",
		"   c c\n" +
			"    c\n" +
			"    c\n" +
			"    :-----.\n" +
			"  .'       `.\n" +
			" / (W)       \\????,\n" +
			"(__,          \\_.'/\n",
		"   c c\n" +
			" ccccccc\n" +
			"    c\n" +
			"    :-----.\n" +
			"  .'       `.\n" +
			" / (W)       \\????,\n" +
			"(__,          \\_.'/\n",
		"   c c\n" +
			" ccccccc\n" +
			" c  c  c\n" +
			"    :-----.\n" +
			"  .'       `.\n" +
			" / (W)       \\????,\n" +
			"(__,          \\_.'/\n",
	},
	"right": []string{
		"                \n" +
			"                \n" +
			"                \n" +
			"        .-----:\n" +
			"      .'       `.\n" +
			",????/       (R) \\\n" +
			"\\`._/          ,__)\n",
		"                \n" +
			"                \n" +
			"                \n" +
			"        .-----:\n" +
			"      .'       `.\n" +
			",????/       (R) \\\n" +
			"\\`._/          ,__)\n",
		"                \n" +
			"                \n" +
			"              c\n" +
			"        .-----:\n" +
			"      .'       `.\n" +
			",????/       (R) \\\n" +
			"\\`._/          ,__)\n",
		"                \n" +
			"              c\n" +
			"              c\n" +
			"        .-----:\n" +
			"      .'       `.\n" +
			",????/       (R) \\\n" +
			"\\`._/          ,__)\n",
		"             c c\n" +
			"              c\n" +
			"              c\n" +
			"        .-----:\n" +
			"      .'       `.\n" +
			",????/       (R) \\\n" +
			"\\`._/          ,__)\n",
		"             c c\n" +
			"           ccccccc\n" +
			"              c\n" +
			"        .-----:\n" +
			"      .'       `.\n" +
			",????/       (R) \\\n" +
			"\\`._/          ,__)\n",
		"             c c\n" +
			"           ccccccc\n" +
			"           c  c  c\n" +
			"        .-----:\n" +
			"      .'       `.\n" +
			",????/       (R) \\\n" +
			"\\`._/          ,__)\n",
	},
}

/*
var snakeShape = termination.Shape {
    "default": []string {
`
                                                          ____
            __??????????????????????????????????????????/   o  \
          /    \????????_?????????????????????_???????/     ____ >
  _??????|  __  |?????/   \????????_????????/   \????|     |
 | \?????|  ||  |????|     |?????/   \?????|     |???|     |
`,
`
                                                          ____
                                             __?????????/   o  \
             _?????????????????????_???????/    \?????/     ____ >
   _???????/   \????????_????????/   \????|  __  |???|     |
  | \?????|     |?????/   \?????|     |???|  ||  |???|     |
`,
`
                                                          ____
                                  __????????????????????/   o  \
 _??????????????????????_???????/    \????????_???????/     ____ >
| \??????????_????????/   \????|  __  |?????/   \????|     |
 \ \???????/   \?????|     |???|  ||  |????|     |???|     |
`,
`
                                                          ____
                       __???????????????????????????????/   o  \
  _??????????_???????/    \????????_??????????????????/     ____ >
 | \???????/   \????|  __  |?????/   \????????_??????|     |
  \ \?????|     |???|  ||  |????|     |?????/   \????|     |
`,
},
}
*/
