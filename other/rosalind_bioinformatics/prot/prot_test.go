package prot

import (
	"testing"
)

func Test_prot(t *testing.T) {
	var cases = []struct {
		fpath string
		protS string
	}{
		{
			"../input/prot/test_input_01.txt",
			"MAMAPRTEINSTRING",
		},
		{
			"../input/prot/test_input_02.txt",
			"MHVLYNVPAFTLPYRRLVRDGKRWFEHPGQETAPRNDLVADPCCREMKALLVIIPQALTYSARGRKCGGRLTPWKEGTTTFRAEHYSLGLSPRYPPDSRCFNAFGSAKRNECRLFPVGAPQWPEKSKTRELLRSSGAIASKSRPWLWRDVVRIWCLDHSNKTLSIECNTIGEIHRTTSSTTRDLRLSFVVIQRGSHDTPPCEASLGLQRDPGHARLHFLSLIANPSYARFLAVPTDPIEVGGGWSLSSHTPTTGSNSYGLLPSQFGHRLSLCNTPCYSAPDFHIRANKPSVTNWLHQGVCRSSSRRCSNVLGTGLYSGYSSRVCLGIYEAFQPFVSTSLALYFFPVTPTINFKHAKILQAYRTCAWFGGCHCKYRITVRRSASLQPTNAWGAPNPLALRKLSVMQMCNTTRHRCYLSFPHCVECRIQCNRHVLSHRLKFIFLLGTPTKGCDMPQSIGITADPPPHRECAFRRGEKNLMQPPVDPGKEPRAGYDREPPCNYIVARTRLIFGNEYAGGIVAVIPLPHVRKLIYIDGISPISYRYPNLSVVRHERSADSISIIHQLYKRARLLRVTMNTKGTSSDCPPERSPELWATQGPRKLSHVFCCYATDKRTTSSMTDRRTSYVKSSPLIFALVSRKTPTSSQIYSTTMLTRVGITPMPIRYFRPAEPGAFAYDVRPNVLYMTAVQGYPYSGRVMAMKRRSFLAERTLRICDSVGDFYKDGLKLYKRSNGGFVDYEENSLAKKSRGTYGDLRPFVRAVQKTWSAFYRFAPVNADVPSMSHYQFIYARRTSQSLRYMNPVIGDSRKQPPVVRLSYSSRQDATYPALTEVDAHCGLPPRFCSVSRSFRSRHSTRYRHCSVRLYGKFSYGRPQAVVRSAGRLRKRLGRRMLTVYFSLQTAIGANRFYKIILRLPARDTCVYLVGSIWLNLLISHDSNRGGMRARNAPSYTYQDKYPKCSEAITRALYKKGQKRIVKRMLLSTRVHGVLDTKSNKLRETSSNQLQHLTNTLSSGKLLFRKKSYKDHTGVGPCPPPLSFIAGHITYAGGPKMAPDNNLCHSIPLVTQANMGLLSVHQVCDRVFGSRQIRYSCTFMRGVETAGCALWEECPSGIFGIWEMTVGVFRALFISIELLLSVAWENSSLPSANCAPLTSIQCELSLSMTSCSLVPVYGGFVSCVSDWFSFLLCLGSETGTGSKLREASGVVLPRVPNSYWPRASGQLIARSLRHALGDTLLAEYVLGGQILITFPGTPLQPVTRHKTASSVLKLRSYRWVLADAVAHPSPLLLHSKRSPPRCYGPYGYASDRLPSRNLALTYLSSAGTTGLGLRPLKRRNLIVLEPVLSTCMQRLHTSCDAVSDSGEWNCQSTKRTTSRLPGRFCQGDLAPLYTKRSDPARPSVALLNSLARDIAAVRDRLRPSRGARVIAIEKVPVLSAGRRSIPRRMPARNGVIRLEVMLSLMTYRNWSVTTTSRKYGMLVFGCPRRIIPLMLLGRVYAFLTTGPPRKIGLHLVVLIGCFSYCITSLRRPQHETCAPLICAVTEPALYIGCTTDGKPGGGRPVLRRMILDSYTLAYLKGLSGKGPPTAHLNVPAYSASRDWSSSGLLTFFKYGREWEETAYAMLGRRMIYAHGAAVRARKPPEAELFGQPRAAFPAVACVKKCKALIQCGTNQHSNRGSPRVLGFLNTATTKSGIGAPTQDCFRGLSRVNDSYQRRVFDAFLSFIMSNCVAPVQFAGVVGHKGRLVVKNQRAPRMTTSYDTGSLSPAKPAVKHKFADFDGLQLFPSQDLKGLTSKDAINCSIPDWFRDASSSSGSAHRLIADPGCDLNGVTTPKCDQHTYLNEVIQRSLQVSRYVIGVLAEFLRAGQSPLSRSIANTYLHFPTMRGRHTTVGDPELLAEHVVEKIKCFREGQSARHEISRRFTQAHEECSSKKPLSKSMVGYLSFIVRVGNNPPNFKRVGDIEVIMRRSPLPRPCPSDCLWFRLDVLTLCCQTNPHMNGPTGSQLAASQVSRGLNDRTGSTKSLRTGHINSAFDRSRGKCFVSLSPSPTYDSGGSDGARTALSQVVNRGSVTRSSPSGFRNKRELTSHNAGSLSSTGGRAVWITPDRAAVISIDPLSETISSSVIESILALVRGTSASPRGSVQSMQFLLFEGRLPSTLVIYNNYAILPQSIGLRVIATLRHRQCVRSLSALLELESAHTSKTSWVLLSALPAKESELLVSKRRRIPHSHLKGSEILYSDSCLSPHFPPVARVPRRGKHRRGWQPISVAVILKLSNYPYMEDYPPVTEYDCQTPVLRTYTPDDMGFSRPVPTPAESGFMLTSELAQVGDTAEALDPNRHEAPRTMPKTTVGRPAVGMAISFGRRNHPVTVLRDECLMHFGPSLTWGITEYSGGLENAAGRSSFECPRRSSVSVIDTIVIEPDGSKPNPTPPIHLLEMLNHQLPSKRFRKGEHIRISIHESLLPCSPLFQSPRTASGSHAVNLTPARWNGFRSLWQLACPLKNLHRTKPPWLKCEVGAITPNRIPGSQTLTGPTAPQGGLISPGVDGGRYLGLPNYPINSLSITKRKASPLKDYRDRIACGIATNVSSSANILFRDLTQPSIARGMDADGLSSTICMAKIEERNQGKRGVTRVLLYHIGEPDAGGPSRCRKSRYGPFRLVHHGARSPIGLCITISDGEVTIRNKGLILKEFRHQQNSRGSLHQLYGIPTGCHRDGRSEFNQWERRNCHSRILHQDRSSRHVRNPHIHSW",
		},
	}
	for _, c := range cases {
		protS, err := prot(c.fpath)
		if err != nil {
			t.Errorf("Error retriving encoded string: %v", err)
		}
		if protS != c.protS {
			t.Errorf("Encoded string does not match - got %v want: %v\n", protS, c.protS)
		}
		//fmt.Printf("file = %v, prob = %v\n", c.fpath, protS)
	}
}
