package twid_test

import (
	"fmt"
	"testing"

	twid "github.com/guanting112/taiwan-id-validator"
	"github.com/stretchr/testify/assert"
)

func Test_Validate_Ok(t *testing.T) {
	rules := []struct {
		description string
		taiwanId    string
		expected    bool
	}{
		{"Valid National ID 1", "A123456789", true},
		{"Valid National ID 2", "Y144766850", true},
		{"Valid National ID (Lowercase)", "a123456789", true},

		{"Invalid National ID Checksum", "A123456788", false},
		{"Invalid National ID Length", "A12345678", false},
		{"Invalid National ID Length 2", "A123456781283130-", false},
		{"Invalid National ID Format", "1123456789", false},
		{"Invalid National ID Format", "@90y8fgtear0g8", false},

		{"Valid New ARC-ID 1", "A824285055", true},
		{"Invalid New ARC-ID 1", "A824285054", false},
		{"Valid New ARC-ID 2", "G951582036", true},
		{"Invalid New ARC-ID 2", "G951582034", false},
		{"Valid New ARC-ID 3", "R884356462", true},
		{"Invalid New ARC-ID 3", "R884356460", false},
		{"Valid New ARC-ID 4", "R884356462", true},
		{"Invalid New ARC-ID 4", "R884356461", false},
		{"Valid New ARC-ID 5", "J839129044", true},
		{"Invalid New ARC-ID 5", "J839129043", false},
		{"Valid New ARC-ID 6", "R884356462", true},
		{"Invalid New ARC-ID 6", "R884356461", false},

		{"Invalid New ARC-ID 1", "A824285058", false},
		{"Invalid New ARC-ID 2", "G951582039", false},
		{"Invalid New ARC-ID 3", "R884356410", false},
		{"Invalid New ARC-ID 4", "R884356425", false},
		{"Invalid New ARC-ID 5", "J839129060", false},
		{"Invalid New ARC-ID 6", "R884356477", false},
		{"Invalid New ARC-ID 7", "A800000000", false},
		{"Invalid New ARC-ID 7", "Z900000000", false},

		{"Invalid ID 1", "9y300tq943hr", false},
		{"Invalid ID 2", "19237420123", false},
		{"Invalid ID 3", "B923742010", false},
		{"Invalid ID 4", "*^!(Y@#G!P@)", false},

		{"Random Id A107033713", "A107033713", true},
		{"Random Id Y118210609", "Y118210609", true},
		{"Random Id E204432327", "E204432327", true},
		{"Random Id T211648211", "T211648211", true},
		{"Random Id C216425009", "C216425009", true},
		{"Random Id Q217525429", "Q217525429", true},
		{"Random Id D110125483", "D110125483", true},
		{"Random Id Y206607630", "Y206607630", true},
		{"Random Id D210346706", "D210346706", true},
		{"Random Id Y110306215", "Y110306215", true},
		{"Random Id Q213382740", "Q213382740", true},
		{"Random Id W102012661", "W102012661", true},
		{"Random Id N200052080", "N200052080", true},
		{"Random Id B114043757", "B114043757", true},
		{"Random Id P210781213", "P210781213", true},
		{"Random Id G217072800", "G217072800", true},
		{"Random Id V203375023", "V203375023", true},
		{"Random Id S107735244", "S107735244", true},
		{"Random Id G115857385", "G115857385", true},
		{"Random Id Y103812128", "Y103812128", true},
		{"Random Id B110158613", "B110158613", true},
		{"Random Id X118488802", "X118488802", true},
		{"Random Id A106023035", "A106023035", true},
		{"Random Id S100128303", "S100128303", true},
		{"Random Id B100344732", "B100344732", true},
		{"Random Id S218216216", "S218216216", true},
		{"Random Id I215752650", "I215752650", true},
		{"Random Id P118137439", "P118137439", true},
		{"Random Id C105035677", "C105035677", true},
		{"Random Id G108537181", "G108537181", true},
		{"Random Id X108318777", "X108318777", true},
		{"Random Id F205066177", "F205066177", true},
		{"Random Id B114357818", "B114357818", true},
		{"Random Id S212652067", "S212652067", true},
		{"Random Id A200266529", "A200266529", true},
		{"Random Id S117012052", "S117012052", true},
		{"Random Id X203175658", "X203175658", true},
		{"Random Id K107268733", "K107268733", true},
		{"Random Id M200871182", "M200871182", true},
		{"Random Id Y105205210", "Y105205210", true},
		{"Random Id W212106672", "W212106672", true},
		{"Random Id S200445025", "S200445025", true},
		{"Random Id O104381760", "O104381760", true},
		{"Random Id T113546781", "T113546781", true},
		{"Random Id U208378774", "U208378774", true},
		{"Random Id T200450036", "T200450036", true},
		{"Random Id O101136227", "O101136227", true},
		{"Random Id F207665221", "F207665221", true},
		{"Random Id F206755511", "F206755511", true},
		{"Random Id I101524164", "I101524164", true},
		{"Random Id X100850741", "X100850741", true},
		{"Random Id T204145483", "T204145483", true},
		{"Random Id L107065081", "L107065081", true},
		{"Random Id C202711547", "C202711547", true},
		{"Random Id N107702830", "N107702830", true},
		{"Random Id A106053382", "A106053382", true},
		{"Random Id X115014319", "X115014319", true},
		{"Random Id Q107250477", "Q107250477", true},
		{"Random Id I204685415", "I204685415", true},
		{"Random Id X203762268", "X203762268", true},
		{"Random Id X105226209", "X105226209", true},
		{"Random Id G110126670", "G110126670", true},
		{"Random Id P104032261", "P104032261", true},
		{"Random Id C215852619", "C215852619", true},
		{"Random Id A103770135", "A103770135", true},
		{"Random Id V114251587", "V114251587", true},
		{"Random Id Q216151318", "Q216151318", true},
		{"Random Id C200700162", "C200700162", true},
		{"Random Id L102662446", "L102662446", true},
		{"Random Id Y116731310", "Y116731310", true},
		{"Random Id C115230620", "C115230620", true},
		{"Random Id O102884786", "O102884786", true},
		{"Random Id A213041205", "A213041205", true},
		{"Random Id C108613806", "C108613806", true},
		{"Random Id S105513722", "S105513722", true},
		{"Random Id N100784852", "N100784852", true},
		{"Random Id E107530659", "E107530659", true},
		{"Random Id U116773441", "U116773441", true},
		{"Random Id D207228235", "D207228235", true},
		{"Random Id P210006742", "P210006742", true},
		{"Random Id W117477156", "W117477156", true},
		{"Random Id K112036725", "K112036725", true},
		{"Random Id Y200654224", "Y200654224", true},
		{"Random Id S100541860", "S100541860", true},
		{"Random Id I113711775", "I113711775", true},
		{"Random Id K108304738", "K108304738", true},
		{"Random Id L115545570", "L115545570", true},
		{"Random Id K214144286", "K214144286", true},
		{"Random Id M114557559", "M114557559", true},
		{"Random Id V203476212", "V203476212", true},
		{"Random Id L210738415", "L210738415", true},
		{"Random Id O211578711", "O211578711", true},
		{"Random Id H114106166", "H114106166", true},
		{"Random Id R201783372", "R201783372", true},
		{"Random Id L118524813", "L118524813", true},
		{"Random Id A101200372", "A101200372", true},
		{"Random Id T107717443", "T107717443", true},
		{"Random Id W116358410", "W116358410", true},
		{"Random Id L102723373", "L102723373", true},
		{"Random Id C102058570", "C102058570", true},
		{"Random Id C105752431", "C105752431", true},
		{"Random Id P213502081", "P213502081", true},
		{"Random Id L107018475", "L107018475", true},
		{"Random Id L108644228", "L108644228", true},
		{"Random Id M204060578", "M204060578", true},
		{"Random Id Y101462344", "Y101462344", true},
		{"Random Id E200706746", "E200706746", true},
		{"Random Id E213085541", "E213085541", true},
		{"Random Id I106712771", "I106712771", true},
		{"Random Id T112713840", "T112713840", true},
		{"Random Id C104168140", "C104168140", true},
		{"Random Id T118754649", "T118754649", true},
		{"Random Id N212071380", "N212071380", true},
		{"Random Id L214485282", "L214485282", true},
		{"Random Id H116780437", "H116780437", true},
		{"Random Id W208444669", "W208444669", true},
		{"Random Id Y212750446", "Y212750446", true},
		{"Random Id J111547737", "J111547737", true},
		{"Random Id D108657434", "D108657434", true},
		{"Random Id P214555522", "P214555522", true},
		{"Random Id M203475584", "M203475584", true},
		{"Random Id P213257232", "P213257232", true},
		{"Random Id G203750579", "G203750579", true},
		{"Random Id O118573669", "O118573669", true},
		{"Random Id X207573525", "X207573525", true},
		{"Random Id Y204332238", "Y204332238", true},
		{"Random Id Q111320115", "Q111320115", true},
		{"Random Id C201656823", "C201656823", true},
		{"Random Id B216541338", "B216541338", true},
		{"Random Id P206785543", "P206785543", true},
		{"Random Id O114755236", "O114755236", true},
		{"Random Id H111123254", "H111123254", true},
		{"Random Id A102236867", "A102236867", true},
		{"Random Id G118825689", "G118825689", true},
		{"Random Id S111662834", "S111662834", true},
		{"Random Id Y110387229", "Y110387229", true},
		{"Random Id I207144840", "I207144840", true},
		{"Random Id D111335438", "D111335438", true},
		{"Random Id L206728072", "L206728072", true},
		{"Random Id V106030449", "V106030449", true},
		{"Random Id M108308602", "M108308602", true},
		{"Random Id O105401467", "O105401467", true},
		{"Random Id K215230629", "K215230629", true},
		{"Random Id P104488227", "P104488227", true},
		{"Random Id Q112820829", "Q112820829", true},
		{"Random Id X205630670", "X205630670", true},
		{"Random Id X217411587", "X217411587", true},
		{"Random Id Q103482080", "Q103482080", true},
		{"Random Id G105838663", "G105838663", true},
		{"Random Id A100380255", "A100380255", true},
		{"Random Id L108712301", "L108712301", true},
		{"Random Id S107827505", "S107827505", true},
		{"Random Id J115272526", "J115272526", true},
		{"Random Id Y101583035", "Y101583035", true},
		{"Random Id A103270649", "A103270649", true},
		{"Random Id M201273402", "M201273402", true},
		{"Random Id U117578628", "U117578628", true},
		{"Random Id Q208477803", "Q208477803", true},
		{"Random Id Q104464577", "Q104464577", true},
		{"Random Id R104380040", "R104380040", true},
		{"Random Id I116736532", "I116736532", true},
		{"Random Id E103124888", "E103124888", true},
		{"Random Id Y114364348", "Y114364348", true},
		{"Random Id I112414848", "I112414848", true},
		{"Random Id W215803276", "W215803276", true},
		{"Random Id M107005620", "M107005620", true},
		{"Random Id R210700676", "R210700676", true},
		{"Random Id C116830464", "C116830464", true},
		{"Random Id S101053770", "S101053770", true},
		{"Random Id V106541634", "V106541634", true},
		{"Random Id U201444668", "U201444668", true},
		{"Random Id C205870445", "C205870445", true},
		{"Random Id Q110053642", "Q110053642", true},
		{"Random Id M203182851", "M203182851", true},
		{"Random Id L101656408", "L101656408", true},
		{"Random Id I103515583", "I103515583", true},
		{"Random Id J104536019", "J104536019", true},
		{"Random Id M201318353", "M201318353", true},
		{"Random Id T101632405", "T101632405", true},
		{"Random Id M208266558", "M208266558", true},
		{"Random Id R202377734", "R202377734", true},
		{"Random Id N203512010", "N203512010", true},
		{"Random Id B103531764", "B103531764", true},
		{"Random Id A107035468", "A107035468", true},
		{"Random Id A208264205", "A208264205", true},
		{"Random Id P200380526", "P200380526", true},
		{"Random Id B107222266", "B107222266", true},
		{"Random Id P218463061", "P218463061", true},
		{"Random Id G214763000", "G214763000", true},
		{"Random Id B208548038", "B208548038", true},
		{"Random Id E111517879", "E111517879", true},
		{"Random Id J213358776", "J213358776", true},
		{"Random Id E217028873", "E217028873", true},
		{"Random Id H111187525", "H111187525", true},
		{"Random Id H108212788", "H108212788", true},
		{"Random Id O218640121", "O218640121", true},
		{"Random Id F118378755", "F118378755", true},
		{"Random Id L208733240", "L208733240", true},
		{"Random Id J113514343", "J113514343", true},
		{"Random Id J114564650", "J114564650", true},

		{"Random Invalid Id", "E207760092", false},
		{"Random Invalid Id", "N211007890", false},
		{"Random Invalid Id", "I203158694", false},
		{"Random Invalid Id", "I103358798", false},
		{"Random Invalid Id", "D207313399", false},
		{"Random Invalid Id", "P200061195", false},
		{"Random Invalid Id", "U104667895", false},
		{"Random Invalid Id", "W211021497", false},
		{"Random Invalid Id", "A115751894", false},
		{"Random Invalid Id", "H213126892", false},
		{"Random Invalid Id", "G201643298", false},
		{"Random Invalid Id", "L206543794", false},
		{"Random Invalid Id", "V116358494", false},
		{"Random Invalid Id", "Q102043891", false},
		{"Random Invalid Id", "L110503193", false},
		{"Random Invalid Id", "L201063696", false},
		{"Random Invalid Id", "Q110413199", false},
		{"Random Invalid Id", "Q204232191", false},
		{"Random Invalid Id", "Y213481894", false},
		{"Random Invalid Id", "O212318893", false},
		{"Random Invalid Id", "A212875794", false},
		{"Random Invalid Id", "Q106371096", false},
		{"Random Invalid Id", "Q115170195", false},
		{"Random Invalid Id", "A217206397", false},
		{"Random Invalid Id", "M112288290", false},
		{"Random Invalid Id", "X101382194", false},
		{"Random Invalid Id", "H103745498", false},
		{"Random Invalid Id", "A207541197", false},
		{"Random Invalid Id", "D116626192", false},
		{"Random Invalid Id", "A214700092", false},
		{"Random Invalid Id", "M100832899", false},
		{"Random Invalid Id", "T115117393", false},
		{"Random Invalid Id", "Q205647290", false},
		{"Random Invalid Id", "T215304592", false},
		{"Random Invalid Id", "J203006796", false},
		{"Random Invalid Id", "X216481794", false},
		{"Random Invalid Id", "S115008595", false},
		{"Random Invalid Id", "X114865099", false},
		{"Random Invalid Id", "J113221699", false},
		{"Random Invalid Id", "A102420296", false},
		{"Random Invalid Id", "Y204147498", false},
		{"Random Invalid Id", "U118542395", false},
		{"Random Invalid Id", "O203320193", false},
		{"Random Invalid Id", "C102541194", false},
		{"Random Invalid Id", "Q116711194", false},
		{"Random Invalid Id", "T200260798", false},
		{"Random Invalid Id", "U202737496", false},
		{"Random Invalid Id", "V114545298", false},
		{"Random Invalid Id", "S201816398", false},
		{"Random Invalid Id", "P215440899", false},
	}

	for _, rule := range rules {
		t.Run(rule.description, func(t *testing.T) {
			assert.Equal(t, rule.expected, twid.Validate(rule.taiwanId))
		})
	}
}

func ExampleValidate() {
	if twid.Validate("A123456789") {
		fmt.Println("Valid ID")
	} else {
		fmt.Println("Invalid ID")
	}

	if twid.Validate("A800000014") {
		fmt.Println("Valid ARC-ID")
	} else {
		fmt.Println("Invalid ARC-ID")
	}

	if twid.Validate("AC01234567") {
		fmt.Println("Valid Old ARC-ID")
	} else {
		fmt.Println("Invalid Old ARC-ID")
	}
}

func Test_ValidateNationId_Ok(t *testing.T) {
	rules := []struct {
		description string
		taiwanId    string
		expected    bool
	}{
		{"Valid National ID 1", "A123456789", true},
		{"Valid National ID 2", "Y144766850", true},
		{"Valid National ID (Lowercase)", "a123456789", true},
		{"Random Id D210346706", "D210346706", true},
		{"Random Id Y110306215", "Y110306215", true},
		{"Random Id Q213382740", "Q213382740", true},
		{"Random Id W102012661", "W102012661", true},
		{"Random Id N200052080", "N200052080", true},
		{"Random Id B114043757", "B114043757", true},
		{"Random Id P210781213", "P210781213", true},
		{"Random Id G217072800", "G217072800", true},
		{"Random Id V203375023", "V203375023", true},
		{"Random Id S107735244", "S107735244", true},
		{"Random Id G115857385", "G115857385", true},
		{"Random Id Y103812128", "Y103812128", true},
		{"Random Invalid Id", "T200260798", false},
		{"Random Invalid Id", "U202737496", false},
		{"Random Invalid Id", "V114545298", false},
		{"Random Invalid Id", "S201816398", false},
		{"Random Invalid Id", "P215440899", false},
	}

	for _, rule := range rules {
		t.Run(rule.description, func(t *testing.T) {
			assert.Equal(t, rule.expected, twid.ValidateNationId(rule.taiwanId))
		})
	}
}

func ExampleValidateNationId() {
	if twid.ValidateNationId("A123456789") {
		fmt.Println("Valid ID")
	} else {
		fmt.Println("Invalid ID")
	}
}

func Test_ValidateArcId_Ok(t *testing.T) {
	rules := []struct {
		description string
		taiwanId    string
		expected    bool
	}{
		{"Valid New ARC-ID 1", "R884356462", true},
		{"Invalid New ARC-ID 1", "R884356461", false},
		{"Valid New ARC-ID 2", "J839129044", true},
		{"Invalid New ARC-ID 2", "J839129043", false},
		{"Valid Old ARC-ID 3", "AC01234567", true},
		{"Invalid Old ARC-ID 1", "AC01234565", false},
		{"Invalid Old ARC-ID 2", "AC01234560", false},
		{"Invalid Old ARC-ID 3", "AC01234561", false},
		{"Invalid Old ARC-ID 4", "AC01234562", false},
	}

	for _, rule := range rules {
		t.Run(rule.description, func(t *testing.T) {
			assert.Equal(t, rule.expected, twid.ValidateArcId(rule.taiwanId))
		})
	}
}

func ExampleValidateArcId() {
	if twid.ValidateArcId("A800000014") {
		fmt.Println("Valid ARC-ID")
	} else {
		fmt.Println("Invalid ARC-ID")
	}

	if twid.ValidateArcId("AC01234567") {
		fmt.Println("Valid Old ARC-ID")
	} else {
		fmt.Println("Invalid Old ARC-ID")
	}
}
