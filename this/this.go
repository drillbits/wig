package wig

import (
	. "github.com/drillbits/wig"
	"io"
	"os"
	"strings"
)

var lyrics = `Gur fpurqhyr’f gvtug ba gur pyhfgre gbavtug
Fb V cnenyyryvmrq zl pbqr

Nyy gubfr guernqf naq pbagvahngvbaf
Zl urnq’f tbvat gb rkcybqr

Naq nyy gung obvyrecyngr
Gung SnpgbelOhvyqreNqncgreQryrtngrVzcy

Frrzf hawhfgvsvrq
Tvir zr fbzrguvat fvzcyr

Qba’g jevgr va Fpurzr,
Qba’g jevgr va P

Ab zber cbvagref
gung V sbetbg gb serr()

Wnin’f ireobfr
Clguba’f gbb fybj

Vg’f gvzr lbh xabj

Jevgr va Tb!
Jevgr va Tb!

Ab vaurevgnapr nalzber

Jevgr va Tb!
Jevgr va Tb!

Gurer’f ab qb be juvyr,
whfg sbe

V qba’g pner jung lbhe yvagref fnl.
V’ir tbg gbbyf sbe gung

Gur pbqr arire obgurerq zr naljnl

Vg’f shaal ubj fbzr srngherf
Znxr rirel punatr frrz fznyy

Naq gur reebef gung bapr fybjrq zr
Qba’g trg zr qbja ng nyy

Vg’f gvzr gb frr jung Tb pna qb
’Pnhfr vg frrzf gbb tbbq gb or gehr

Ab ybat pbzcvyr gvzrf sbe zr.
V’z serr!

Jevgr va Tb!
Jevgr va Tb!

Xvff lbhe cbvagre zngu tbbqolr

Jevgr va Tb!
Jevgr va Tb!

Gvzr gb tvir TP n gel

V qba’g pner vs zl fgehpgherf fgnl
Ba gur urnc be fgnpx

// GBQB(fpnynovyvgl): zhfvpny vagreyhqr

Zl cebtenz fcnjaf vgf tbebhgvarf
jvgubhg n fbhaq

Pbageby vf fcvenyvat guebhtu
ohssrerq punaaryf nyy nebhaq

V qba’g erzrzore jul V rire bapr fhopynffrq

V’z arire tbvat onpx

Zl grfgf nyy ohvyq naq cnff

Jevgr va Tb!
Jevgr va Tb!

Lbh jba’g hfr Rpyvcfr nalzber

Jevgr va Tb!
Jevgr va Tb!

Jub pnerf jung Obbfg vf sbe?

V qba’g pner jung gur grpu yrnqf fnl
V’yy erjevgr vg nyy!

Jevgvat pbqr arire obgurerq zr, naljnl.
`

func init() {
	s := strings.NewReader(lyrics)
	r := NewInterpretReader(s)
	// r := NewLookupReader(s)
	io.Copy(os.Stdout, &r)
}
