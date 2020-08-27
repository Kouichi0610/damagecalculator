/*
	和名->ID(英名)変換
*/
package index

import (
	"fmt"
)

type pokeIndex struct {
	en string // 英語名
	jp string // 日本語名
}

func JName(name string) (string, error) {
	for _, d := range indices {
		en := d.en
		if name == en {
			return d.jp, nil
		}
	}
	return "", fmt.Errorf("%s not found.", name)
}

// 英名または和名をIDに変換
func Index(name string) (string, error) {
	for _, d := range indices {
		en := d.en
		jp := d.jp
		if name == en || name == jp {
			return d.en, nil
		}
	}
	return "", fmt.Errorf("%s not found.", name)
}

func IndexArray() []string {
	res := make([]string, 0)
	for _, i := range indices {
		res = append(res, i.en)
	}
	return res
}
func JpNamesArray() []string {
	res := make([]string, 0)
	for _, i := range indices {
		res = append(res, i.jp)
	}
	return res
}

// TODO:pokeapiからリスト取れそうなら
var indices []*pokeIndex = []*pokeIndex{
	{"bulbasaur", "フシギダネ"},
	{"ivysaur", "フシギソウ"},
	{"venusaur", "フシギバナ"},
	{"charmander", "ヒトカゲ"},
	{"charmeleon", "リザード"},
	{"charizard", "リザードン"},
	{"squirtle", "ゼニガメ"},
	{"wartortle", "カメール"},
	{"blastoise", "カメックス"},
	{"caterpie", "キャタピー"},
	{"metapod", "トランセル"},
	{"butterfree", "バタフリー"},
	{"weedle", "ビードル"},
	{"kakuna", "コクーン"},
	{"beedrill", "スピアー"},
	{"pidgey", "ポッポ"},
	{"pidgeotto", "ピジョン"},
	{"pidgeot", "ピジョット"},
	{"rattata", "コラッタ"},
	{"raticate", "ラッタ"},
	{"spearow", "オニスズメ"},
	{"fearow", "オニドリル"},
	{"ekans", "アーボ"},
	{"arbok", "アーボック"},
	{"pikachu", "ピカチュウ"},
	{"raichu", "ライチュウ"},
	{"sandshrew", "サンド"},
	{"sandslash", "サンドパン"},
	{"nidoran-f", "ニドラン♀"},
	{"nidorina", "ニドリーナ"},
	{"nidoqueen", "ニドクイン"},
	{"nidoran-m", "ニドラン♂"},
	{"nidorino", "ニドリーノ"},
	{"nidoking", "ニドキング"},
	{"clefairy", "ピッピ"},
	{"clefable", "ピクシー"},
	{"vulpix", "ロコン"},
	{"ninetales", "キュウコン"},
	{"jigglypuff", "プリン"},
	{"wigglytuff", "プクリン"},
	{"zubat", "ズバット"},
	{"golbat", "ゴルバット"},
	{"oddish", "ナゾノクサ"},
	{"gloom", "クサイハナ"},
	{"vileplume", "ラフレシア"},
	{"paras", "パラス"},
	{"parasect", "パラセクト"},
	{"venonat", "コンパン"},
	{"venomoth", "モルフォン"},
	{"diglett", "ディグダ"},
	{"dugtrio", "ダグトリオ"},
	{"meowth", "ニャース"},
	{"persian", "ペルシアン"},
	{"psyduck", "コダック"},
	{"golduck", "ゴルダック"},
	{"mankey", "マンキー"},
	{"primeape", "オコリザル"},
	{"growlithe", "ガーディ"},
	{"arcanine", "ウインディ"},
	{"poliwag", "ニョロモ"},
	{"poliwhirl", "ニョロゾ"},
	{"poliwrath", "ニョロボン"},
	{"abra", "ケーシィ"},
	{"kadabra", "ユンゲラー"},
	{"alakazam", "フーディン"},
	{"machop", "ワンリキー"},
	{"machoke", "ゴーリキー"},
	{"machamp", "カイリキー"},
	{"bellsprout", "マダツボミ"},
	{"weepinbell", "ウツドン"},
	{"victreebel", "ウツボット"},
	{"tentacool", "メノクラゲ"},
	{"tentacruel", "ドククラゲ"},
	{"geodude", "イシツブテ"},
	{"graveler", "ゴローン"},
	{"golem", "ゴローニャ"},
	{"ponyta", "ポニータ"},
	{"rapidash", "ギャロップ"},
	{"slowpoke", "ヤドン"},
	{"slowbro", "ヤドラン"},
	{"magnemite", "コイル"},
	{"magneton", "レアコイル"},
	{"farfetchd", "カモネギ"},
	{"doduo", "ドードー"},
	{"dodrio", "ドードリオ"},
	{"seel", "パウワウ"},
	{"dewgong", "ジュゴン"},
	{"grimer", "ベトベター"},
	{"muk", "ベトベトン"},
	{"shellder", "シェルダー"},
	{"cloyster", "パルシェン"},
	{"gastly", "ゴース"},
	{"haunter", "ゴースト"},
	{"gengar", "ゲンガー"},
	{"onix", "イワーク"},
	{"drowzee", "スリープ"},
	{"hypno", "スリーパー"},
	{"krabby", "クラブ"},
	{"kingler", "キングラー"},
	{"voltorb", "ビリリダマ"},
	{"electrode", "マルマイン"},
	{"exeggcute", "タマタマ"},
	{"exeggutor", "ナッシー"},
	{"cubone", "カラカラ"},
	{"marowak", "ガラガラ"},
	{"hitmonlee", "サワムラー"},
	{"hitmonchan", "エビワラー"},
	{"lickitung", "ベロリンガ"},
	{"koffing", "ドガース"},
	{"weezing", "マタドガス"},
	{"rhyhorn", "サイホーン"},
	{"rhydon", "サイドン"},
	{"chansey", "ラッキー"},
	{"tangela", "モンジャラ"},
	{"kangaskhan", "ガルーラ"},
	{"horsea", "タッツー"},
	{"seadra", "シードラ"},
	{"goldeen", "トサキント"},
	{"seaking", "アズマオウ"},
	{"staryu", "ヒトデマン"},
	{"starmie", "スターミー"},
	{"mr-mime", "バリヤード"},
	{"scyther", "ストライク"},
	{"jynx", "ルージュラ"},
	{"electabuzz", "エレブー"},
	{"magmar", "ブーバー"},
	{"pinsir", "カイロス"},
	{"tauros", "ケンタロス"},
	{"magikarp", "コイキング"},
	{"gyarados", "ギャラドス"},
	{"lapras", "ラプラス"},
	{"ditto", "メタモン"},
	{"eevee", "イーブイ"},
	{"vaporeon", "シャワーズ"},
	{"jolteon", "サンダース"},
	{"flareon", "ブースター"},
	{"porygon", "ポリゴン"},
	{"omanyte", "オムナイト"},
	{"omastar", "オムスター"},
	{"kabuto", "カブト"},
	{"kabutops", "カブトプス"},
	{"aerodactyl", "プテラ"},
	{"snorlax", "カビゴン"},
	{"articuno", "フリーザー"},
	{"zapdos", "サンダー"},
	{"moltres", "ファイヤー"},
	{"dratini", "ミニリュウ"},
	{"dragonair", "ハクリュー"},
	{"dragonite", "カイリュー"},
	{"mewtwo", "ミュウツー"},
	{"mew", "ミュウ"},
	{"chikorita", "チコリータ"},
	{"bayleef", "ベイリーフ"},
	{"meganium", "メガニウム"},
	{"cyndaquil", "ヒノアラシ"},
	{"quilava", "マグマラシ"},
	{"typhlosion", "バクフーン"},
	{"totodile", "ワニノコ"},
	{"croconaw", "アリゲイツ"},
	{"feraligatr", "オーダイル"},
	{"sentret", "オタチ"},
	{"furret", "オオタチ"},
	{"hoothoot", "ホーホー"},
	{"noctowl", "ヨルノズク"},
	{"ledyba", "レディバ"},
	{"ledian", "レディアン"},
	{"spinarak", "イトマル"},
	{"ariados", "アリアドス"},
	{"crobat", "クロバット"},
	{"chinchou", "チョンチー"},
	{"lanturn", "ランターン"},
	{"pichu", "ピチュー"},
	{"cleffa", "ピィ"},
	{"igglybuff", "ププリン"},
	{"togepi", "トゲピー"},
	{"togetic", "トゲチック"},
	{"natu", "ネイティ"},
	{"xatu", "ネイティオ"},
	{"mareep", "メリープ"},
	{"flaaffy", "モココ"},
	{"ampharos", "デンリュウ"},
	{"bellossom", "キレイハナ"},
	{"marill", "マリル"},
	{"azumarill", "マリルリ"},
	{"sudowoodo", "ウソッキー"},
	{"politoed", "ニョロトノ"},
	{"hoppip", "ハネッコ"},
	{"skiploom", "ポポッコ"},
	{"jumpluff", "ワタッコ"},
	{"aipom", "エイパム"},
	{"sunkern", "ヒマナッツ"},
	{"sunflora", "キマワリ"},
	{"yanma", "ヤンヤンマ"},
	{"wooper", "ウパー"},
	{"quagsire", "ヌオー"},
	{"espeon", "エーフィ"},
	{"umbreon", "ブラッキー"},
	{"murkrow", "ヤミカラス"},
	{"slowking", "ヤドキング"},
	{"misdreavus", "ムウマ"},
	{"unown", "アンノーン"},
	{"wobbuffet", "ソーナンス"},
	{"girafarig", "キリンリキ"},
	{"pineco", "クヌギダマ"},
	{"forretress", "フォレトス"},
	{"dunsparce", "ノコッチ"},
	{"gligar", "グライガー"},
	{"steelix", "ハガネール"},
	{"snubbull", "ブルー"},
	{"granbull", "グランブル"},
	{"qwilfish", "ハリーセン"},
	{"scizor", "ハッサム"},
	{"shuckle", "ツボツボ"},
	{"heracross", "ヘラクロス"},
	{"sneasel", "ニューラ"},
	{"teddiursa", "ヒメグマ"},
	{"ursaring", "リングマ"},
	{"slugma", "マグマッグ"},
	{"magcargo", "マグカルゴ"},
	{"swinub", "ウリムー"},
	{"piloswine", "イノムー"},
	{"corsola", "サニーゴ"},
	{"remoraid", "テッポウオ"},
	{"octillery", "オクタン"},
	{"delibird", "デリバード"},
	{"mantine", "マンタイン"},
	{"skarmory", "エアームド"},
	{"houndour", "デルビル"},
	{"houndoom", "ヘルガー"},
	{"kingdra", "キングドラ"},
	{"phanpy", "ゴマゾウ"},
	{"donphan", "ドンファン"},
	{"porygon2", "ポリゴン２"},
	{"stantler", "オドシシ"},
	{"smeargle", "ドーブル"},
	{"tyrogue", "バルキー"},
	{"hitmontop", "カポエラー"},
	{"smoochum", "ムチュール"},
	{"elekid", "エレキッド"},
	{"magby", "ブビィ"},
	{"miltank", "ミルタンク"},
	{"blissey", "ハピナス"},
	{"raikou", "ライコウ"},
	{"entei", "エンテイ"},
	{"suicune", "スイクン"},
	{"larvitar", "ヨーギラス"},
	{"pupitar", "サナギラス"},
	{"tyranitar", "バンギラス"},
	{"lugia", "ルギア"},
	{"ho-oh", "ホウオウ"},
	{"celebi", "セレビィ"},
	{"treecko", "キモリ"},
	{"grovyle", "ジュプトル"},
	{"sceptile", "ジュカイン"},
	{"torchic", "アチャモ"},
	{"combusken", "ワカシャモ"},
	{"blaziken", "バシャーモ"},
	{"mudkip", "ミズゴロウ"},
	{"marshtomp", "ヌマクロー"},
	{"swampert", "ラグラージ"},
	{"poochyena", "ポチエナ"},
	{"mightyena", "グラエナ"},
	{"zigzagoon", "ジグザグマ"},
	{"linoone", "マッスグマ"},
	{"wurmple", "ケムッソ"},
	{"silcoon", "カラサリス"},
	{"beautifly", "アゲハント"},
	{"cascoon", "マユルド"},
	{"dustox", "ドクケイル"},
	{"lotad", "ハスボー"},
	{"lombre", "ハスブレロ"},
	{"ludicolo", "ルンパッパ"},
	{"seedot", "タネボー"},
	{"nuzleaf", "コノハナ"},
	{"shiftry", "ダーテング"},
	{"taillow", "スバメ"},
	{"swellow", "オオスバメ"},
	{"wingull", "キャモメ"},
	{"pelipper", "ペリッパー"},
	{"ralts", "ラルトス"},
	{"kirlia", "キルリア"},
	{"gardevoir", "サーナイト"},
	{"surskit", "アメタマ"},
	{"masquerain", "アメモース"},
	{"shroomish", "キノココ"},
	{"breloom", "キノガッサ"},
	{"slakoth", "ナマケロ"},
	{"vigoroth", "ヤルキモノ"},
	{"slaking", "ケッキング"},
	{"nincada", "ツチニン"},
	{"ninjask", "テッカニン"},
	{"shedinja", "ヌケニン"},
	{"whismur", "ゴニョニョ"},
	{"loudred", "ドゴーム"},
	{"exploud", "バクオング"},
	{"makuhita", "マクノシタ"},
	{"hariyama", "ハリテヤマ"},
	{"azurill", "ルリリ"},
	{"nosepass", "ノズパス"},
	{"skitty", "エネコ"},
	{"delcatty", "エネコロロ"},
	{"sableye", "ヤミラミ"},
	{"mawile", "クチート"},
	{"aron", "ココドラ"},
	{"lairon", "コドラ"},
	{"aggron", "ボスゴドラ"},
	{"meditite", "アサナン"},
	{"medicham", "チャーレム"},
	{"electrike", "ラクライ"},
	{"manectric", "ライボルト"},
	{"plusle", "プラスル"},
	{"minun", "マイナン"},
	{"volbeat", "バルビート"},
	{"illumise", "イルミーゼ"},
	{"roselia", "ロゼリア"},
	{"gulpin", "ゴクリン"},
	{"swalot", "マルノーム"},
	{"carvanha", "キバニア"},
	{"sharpedo", "サメハダー"},
	{"wailmer", "ホエルコ"},
	{"wailord", "ホエルオー"},
	{"numel", "ドンメル"},
	{"camerupt", "バクーダ"},
	{"torkoal", "コータス"},
	{"spoink", "バネブー"},
	{"grumpig", "ブーピッグ"},
	{"spinda", "パッチール"},
	{"trapinch", "ナックラー"},
	{"vibrava", "ビブラーバ"},
	{"flygon", "フライゴン"},
	{"cacnea", "サボネア"},
	{"cacturne", "ノクタス"},
	{"swablu", "チルット"},
	{"altaria", "チルタリス"},
	{"zangoose", "ザングース"},
	{"seviper", "ハブネーク"},
	{"lunatone", "ルナトーン"},
	{"solrock", "ソルロック"},
	{"barboach", "ドジョッチ"},
	{"whiscash", "ナマズン"},
	{"corphish", "ヘイガニ"},
	{"crawdaunt", "シザリガー"},
	{"baltoy", "ヤジロン"},
	{"claydol", "ネンドール"},
	{"lileep", "リリーラ"},
	{"cradily", "ユレイドル"},
	{"anorith", "アノプス"},
	{"armaldo", "アーマルド"},
	{"feebas", "ヒンバス"},
	{"milotic", "ミロカロス"},
	{"castform", "ポワルン"},
	{"kecleon", "カクレオン"},
	{"shuppet", "カゲボウズ"},
	{"banette", "ジュペッタ"},
	{"duskull", "ヨマワル"},
	{"dusclops", "サマヨール"},
	{"tropius", "トロピウス"},
	{"chimecho", "チリーン"},
	{"absol", "アブソル"},
	{"wynaut", "ソーナノ"},
	{"snorunt", "ユキワラシ"},
	{"glalie", "オニゴーリ"},
	{"spheal", "タマザラシ"},
	{"sealeo", "トドグラー"},
	{"walrein", "トドゼルガ"},
	{"clamperl", "パールル"},
	{"huntail", "ハンテール"},
	{"gorebyss", "サクラビス"},
	{"relicanth", "ジーランス"},
	{"luvdisc", "ラブカス"},
	{"bagon", "タツベイ"},
	{"shelgon", "コモルー"},
	{"salamence", "ボーマンダ"},
	{"beldum", "ダンバル"},
	{"metang", "メタング"},
	{"metagross", "メタグロス"},
	{"regirock", "レジロック"},
	{"regice", "レジアイス"},
	{"registeel", "レジスチル"},
	{"latias", "ラティアス"},
	{"latios", "ラティオス"},
	{"kyogre", "カイオーガ"},
	{"groudon", "グラードン"},
	{"rayquaza", "レックウザ"},
	{"jirachi", "ジラーチ"},
	{"deoxys-normal", "デオキシス"},
	{"turtwig", "ナエトル"},
	{"grotle", "ハヤシガメ"},
	{"torterra", "ドダイトス"},
	{"chimchar", "ヒコザル"},
	{"monferno", "モウカザル"},
	{"infernape", "ゴウカザル"},
	{"piplup", "ポッチャマ"},
	{"prinplup", "ポッタイシ"},
	{"empoleon", "エンペルト"},
	{"starly", "ムックル"},
	{"staravia", "ムクバード"},
	{"staraptor", "ムクホーク"},
	{"bidoof", "ビッパ"},
	{"bibarel", "ビーダル"},
	{"kricketot", "コロボーシ"},
	{"kricketune", "コロトック"},
	{"shinx", "コリンク"},
	{"luxio", "ルクシオ"},
	{"luxray", "レントラー"},
	{"budew", "スボミー"},
	{"roserade", "ロズレイド"},
	{"cranidos", "ズガイドス"},
	{"rampardos", "ラムパルド"},
	{"shieldon", "タテトプス"},
	{"bastiodon", "トリデプス"},
	{"burmy", "ミノムッチ"},
	{"wormadam-plant", "ミノマダム(くさきのミノ)"},
	{"mothim", "ガーメイル"},
	{"combee", "ミツハニー"},
	{"vespiquen", "ビークイン"},
	{"pachirisu", "パチリス"},
	{"buizel", "ブイゼル"},
	{"floatzel", "フローゼル"},
	{"cherubi", "チェリンボ"},
	{"cherrim", "チェリム"},
	{"shellos", "カラナクシ"},
	{"gastrodon", "トリトドン"},
	{"ambipom", "エテボース"},
	{"drifloon", "フワンテ"},
	{"drifblim", "フワライド"},
	{"buneary", "ミミロル"},
	{"lopunny", "ミミロップ"},
	{"mismagius", "ムウマージ"},
	{"honchkrow", "ドンカラス"},
	{"glameow", "ニャルマー"},
	{"purugly", "ブニャット"},
	{"chingling", "リーシャン"},
	{"stunky", "スカンプー"},
	{"skuntank", "スカタンク"},
	{"bronzor", "ドーミラー"},
	{"bronzong", "ドータクン"},
	{"bonsly", "ウソハチ"},
	{"mime-jr", "マネネ"},
	{"happiny", "ピンプク"},
	{"chatot", "ペラップ"},
	{"spiritomb", "ミカルゲ"},
	{"gible", "フカマル"},
	{"gabite", "ガバイト"},
	{"garchomp", "ガブリアス"},
	{"munchlax", "ゴンベ"},
	{"riolu", "リオル"},
	{"lucario", "ルカリオ"},
	{"hippopotas", "ヒポポタス"},
	{"hippowdon", "カバルドン"},
	{"skorupi", "スコルピ"},
	{"drapion", "ドラピオン"},
	{"croagunk", "グレッグル"},
	{"toxicroak", "ドクロッグ"},
	{"carnivine", "マスキッパ"},
	{"finneon", "ケイコウオ"},
	{"lumineon", "ネオラント"},
	{"mantyke", "タマンタ"},
	{"snover", "ユキカブリ"},
	{"abomasnow", "ユキノオー"},
	{"weavile", "マニューラ"},
	{"magnezone", "ジバコイル"},
	{"lickilicky", "ベロベルト"},
	{"rhyperior", "ドサイドン"},
	{"tangrowth", "モジャンボ"},
	{"electivire", "エレキブル"},
	{"magmortar", "ブーバーン"},
	{"togekiss", "トゲキッス"},
	{"yanmega", "メガヤンマ"},
	{"leafeon", "リーフィア"},
	{"glaceon", "グレイシア"},
	{"gliscor", "グライオン"},
	{"mamoswine", "マンムー"},
	{"porygon-z", "ポリゴンＺ"},
	{"gallade", "エルレイド"},
	{"probopass", "ダイノーズ"},
	{"dusknoir", "ヨノワール"},
	{"froslass", "ユキメノコ"},
	{"rotom", "ロトム"},
	{"uxie", "ユクシー"},
	{"mesprit", "エムリット"},
	{"azelf", "アグノム"},
	{"dialga", "ディアルガ"},
	{"palkia", "パルキア"},
	{"heatran", "ヒードラン"},
	{"regigigas", "レジギガス"},
	{"giratina-altered", "ギラティナ"},
	{"cresselia", "クレセリア"},
	{"phione", "フィオネ"},
	{"manaphy", "マナフィ"},
	{"darkrai", "ダークライ"},
	{"shaymin-land", "シェイミ"},
	{"arceus", "アルセウス"},
	{"victini", "ビクティニ"},
	{"snivy", "ツタージャ"},
	{"servine", "ジャノビー"},
	{"serperior", "ジャローダ"},
	{"tepig", "ポカブ"},
	{"pignite", "チャオブー"},
	{"emboar", "エンブオー"},
	{"oshawott", "ミジュマル"},
	{"dewott", "フタチマル"},
	{"samurott", "ダイケンキ"},
	{"patrat", "ミネズミ"},
	{"watchog", "ミルホッグ"},
	{"lillipup", "ヨーテリー"},
	{"herdier", "ハーデリア"},
	{"stoutland", "ムーランド"},
	{"purrloin", "チョロネコ"},
	{"liepard", "レパルダス"},
	{"pansage", "ヤナップ"},
	{"simisage", "ヤナッキー"},
	{"pansear", "バオップ"},
	{"simisear", "バオッキー"},
	{"panpour", "ヒヤップ"},
	{"simipour", "ヒヤッキー"},
	{"munna", "ムンナ"},
	{"musharna", "ムシャーナ"},
	{"pidove", "マメパト"},
	{"tranquill", "ハトーボー"},
	{"unfezant", "ケンホロウ"},
	{"blitzle", "シママ"},
	{"zebstrika", "ゼブライカ"},
	{"roggenrola", "ダンゴロ"},
	{"boldore", "ガントル"},
	{"gigalith", "ギガイアス"},
	{"woobat", "コロモリ"},
	{"swoobat", "ココロモリ"},
	{"drilbur", "モグリュー"},
	{"excadrill", "ドリュウズ"},
	{"audino", "タブンネ"},
	{"timburr", "ドッコラー"},
	{"gurdurr", "ドテッコツ"},
	{"conkeldurr", "ローブシン"},
	{"tympole", "オタマロ"},
	{"palpitoad", "ガマガル"},
	{"seismitoad", "ガマゲロゲ"},
	{"throh", "ナゲキ"},
	{"sawk", "ダゲキ"},
	{"sewaddle", "クルミル"},
	{"swadloon", "クルマユ"},
	{"leavanny", "ハハコモリ"},
	{"venipede", "フシデ"},
	{"whirlipede", "ホイーガ"},
	{"scolipede", "ペンドラー"},
	{"cottonee", "モンメン"},
	{"whimsicott", "エルフーン"},
	{"petilil", "チュリネ"},
	{"lilligant", "ドレディア"},
	{"basculin-red-striped", "バスラオ"},
	{"sandile", "メグロコ"},
	{"krokorok", "ワルビル"},
	{"krookodile", "ワルビアル"},
	{"darumaka", "ダルマッカ"},
	{"darmanitan-standard", "ヒヒダルマ"},
	{"maractus", "マラカッチ"},
	{"dwebble", "イシズマイ"},
	{"crustle", "イワパレス"},
	{"scraggy", "ズルッグ"},
	{"scrafty", "ズルズキン"},
	{"sigilyph", "シンボラー"},
	{"yamask", "デスマス"},
	{"cofagrigus", "デスカーン"},
	{"tirtouga", "プロトーガ"},
	{"carracosta", "アバゴーラ"},
	{"archen", "アーケン"},
	{"archeops", "アーケオス"},
	{"trubbish", "ヤブクロン"},
	{"garbodor", "ダストダス"},
	{"zorua", "ゾロア"},
	{"zoroark", "ゾロアーク"},
	{"minccino", "チラーミィ"},
	{"cinccino", "チラチーノ"},
	{"gothita", "ゴチム"},
	{"gothorita", "ゴチミル"},
	{"gothitelle", "ゴチルゼル"},
	{"solosis", "ユニラン"},
	{"duosion", "ダブラン"},
	{"reuniclus", "ランクルス"},
	{"ducklett", "コアルヒー"},
	{"swanna", "スワンナ"},
	{"vanillite", "バニプッチ"},
	{"vanillish", "バニリッチ"},
	{"vanilluxe", "バイバニラ"},
	{"deerling", "シキジカ"},
	{"sawsbuck", "メブキジカ"},
	{"emolga", "エモンガ"},
	{"karrablast", "カブルモ"},
	{"escavalier", "シュバルゴ"},
	{"foongus", "タマゲタケ"},
	{"amoonguss", "モロバレル"},
	{"frillish", "プルリル"},
	{"jellicent", "ブルンゲル"},
	{"alomomola", "ママンボウ"},
	{"joltik", "バチュル"},
	{"galvantula", "デンチュラ"},
	{"ferroseed", "テッシード"},
	{"ferrothorn", "ナットレイ"},
	{"klink", "ギアル"},
	{"klang", "ギギアル"},
	{"klinklang", "ギギギアル"},
	{"tynamo", "シビシラス"},
	{"eelektrik", "シビビール"},
	{"eelektross", "シビルドン"},
	{"elgyem", "リグレー"},
	{"beheeyem", "オーベム"},
	{"litwick", "ヒトモシ"},
	{"lampent", "ランプラー"},
	{"chandelure", "シャンデラ"},
	{"axew", "キバゴ"},
	{"fraxure", "オノンド"},
	{"haxorus", "オノノクス"},
	{"cubchoo", "クマシュン"},
	{"beartic", "ツンベアー"},
	{"cryogonal", "フリージオ"},
	{"shelmet", "チョボマキ"},
	{"accelgor", "アギルダー"},
	{"stunfisk", "マッギョ"},
	{"mienfoo", "コジョフー"},
	{"mienshao", "コジョンド"},
	{"druddigon", "クリムガン"},
	{"golett", "ゴビット"},
	{"golurk", "ゴルーグ"},
	{"pawniard", "コマタナ"},
	{"bisharp", "キリキザン"},
	{"bouffalant", "バッフロン"},
	{"rufflet", "ワシボン"},
	{"braviary", "ウォーグル"},
	{"vullaby", "バルチャイ"},
	{"mandibuzz", "バルジーナ"},
	{"heatmor", "クイタラン"},
	{"durant", "アイアント"},
	{"deino", "モノズ"},
	{"zweilous", "ジヘッド"},
	{"hydreigon", "サザンドラ"},
	{"larvesta", "メラルバ"},
	{"volcarona", "ウルガモス"},
	{"cobalion", "コバルオン"},
	{"terrakion", "テラキオン"},
	{"virizion", "ビリジオン"},
	{"tornadus-incarnate", "トルネロス"},
	{"thundurus-incarnate", "ボルトロス"},
	{"reshiram", "レシラム"},
	{"zekrom", "ゼクロム"},
	{"landorus-incarnate", "ランドロス"},
	{"kyurem", "キュレム"},
	{"keldeo-ordinary", "ケルディオ"},
	{"meloetta-aria", "メロエッタ"},
	{"genesect", "ゲノセクト"},
	{"chespin", "ハリマロン"},
	{"quilladin", "ハリボーグ"},
	{"chesnaught", "ブリガロン"},
	{"fennekin", "フォッコ"},
	{"braixen", "テールナー"},
	{"delphox", "マフォクシー"},
	{"froakie", "ケロマツ"},
	{"frogadier", "ゲコガシラ"},
	{"greninja", "ゲッコウガ"},
	{"bunnelby", "ホルビー"},
	{"diggersby", "ホルード"},
	{"fletchling", "ヤヤコマ"},
	{"fletchinder", "ヒノヤコマ"},
	{"talonflame", "ファイアロー"},
	{"scatterbug", "コフキムシ"},
	{"spewpa", "コフーライ"},
	{"vivillon", "ビビヨン"},
	{"litleo", "シシコ"},
	{"pyroar", "カエンジシ"},
	{"flabebe", "フラベベ"},
	{"floette", "フラエッテ"},
	{"florges", "フラージェス"},
	{"skiddo", "メェークル"},
	{"gogoat", "ゴーゴート"},
	{"pancham", "ヤンチャム"},
	{"pangoro", "ゴロンダ"},
	{"furfrou", "トリミアン"},
	{"espurr", "ニャスパー"},
	{"meowstic-male", "ニャオニクス♂"},
	{"honedge", "ヒトツキ"},
	{"doublade", "ニダンギル"},
	{"aegislash-shield", "ギルガルド(シールド)"},
	{"spritzee", "シュシュプ"},
	{"aromatisse", "フレフワン"},
	{"swirlix", "ペロッパフ"},
	{"slurpuff", "ペロリーム"},
	{"inkay", "マーイーカ"},
	{"malamar", "カラマネロ"},
	{"binacle", "カメテテ"},
	{"barbaracle", "ガメノデス"},
	{"skrelp", "クズモー"},
	{"dragalge", "ドラミドロ"},
	{"clauncher", "ウデッポウ"},
	{"clawitzer", "ブロスター"},
	{"helioptile", "エリキテル"},
	{"heliolisk", "エレザード"},
	{"tyrunt", "チゴラス"},
	{"tyrantrum", "ガチゴラス"},
	{"amaura", "アマルス"},
	{"aurorus", "アマルルガ"},
	{"sylveon", "ニンフィア"},
	{"hawlucha", "ルチャブル"},
	{"dedenne", "デデンネ"},
	{"carbink", "メレシー"},
	{"goomy", "ヌメラ"},
	{"sliggoo", "ヌメイル"},
	{"goodra", "ヌメルゴン"},
	{"klefki", "クレッフィ"},
	{"phantump", "ボクレー"},
	{"trevenant", "オーロット"},
	{"pumpkaboo-average", "バケッチャ平均"},
	{"gourgeist-average", "パンプジン平均"},
	{"bergmite", "カチコール"},
	{"avalugg", "クレベース"},
	{"noibat", "オンバット"},
	{"noivern", "オンバーン"},
	{"xerneas", "ゼルネアス"},
	{"yveltal", "イベルタル"},
	{"zygarde", "ジガルデ"},
	{"diancie", "ディアンシー"},
	{"hoopa", "フーパ"},
	{"volcanion", "ボルケニオン"},
	{"rowlet", "モクロー"},
	{"dartrix", "フクスロー"},
	{"decidueye", "ジュナイパー"},
	{"litten", "ニャビー"},
	{"torracat", "ニャヒート"},
	{"incineroar", "ガオガエン"},
	{"popplio", "アシマリ"},
	{"brionne", "オシャマリ"},
	{"primarina", "アシレーヌ"},
	{"pikipek", "ツツケラ"},
	{"trumbeak", "ケララッパ"},
	{"toucannon", "ドデカバシ"},
	{"yungoos", "ヤングース"},
	{"gumshoos", "デカグース"},
	{"grubbin", "アゴジムシ"},
	{"charjabug", "デンヂムシ"},
	{"vikavolt", "クワガノン"},
	{"crabrawler", "マケンカニ"},
	{"crabominable", "ケケンカニ"},
	{"oricorio-baile", "オドリドリめらめら"},
	{"cutiefly", "アブリー"},
	{"ribombee", "アブリボン"},
	{"rockruff", "イワンコ"},
	{"lycanroc-midday", "ルガルガン(まひる)"},
	{"wishiwashi-solo", "ヨワシ(単独)"},
	{"mareanie", "ヒドイデ"},
	{"toxapex", "ドヒドイデ"},
	{"mudbray", "ドロバンコ"},
	{"mudsdale", "バンバドロ"},
	{"dewpider", "シズクモ"},
	{"araquanid", "オニシズクモ"},
	{"fomantis", "カリキリ"},
	{"lurantis", "ラランテス"},
	{"morelull", "ネマシュ"},
	{"shiinotic", "マシェード"},
	{"salandit", "ヤトウモリ"},
	{"salazzle", "エンニュート"},
	{"stufful", "ヌイコグマ"},
	{"bewear", "キテルグマ"},
	{"bounsweet", "アマカジ"},
	{"steenee", "アママイコ"},
	{"tsareena", "アマージョ"},
	{"comfey", "キュワワー"},
	{"oranguru", "ヤレユータン"},
	{"passimian", "ナゲツケサル"},
	{"wimpod", "コソクムシ"},
	{"golisopod", "グソクムシャ"},
	{"sandygast", "スナバァ"},
	{"palossand", "シロデスナ"},
	{"pyukumuku", "ナマコブシ"},
	{"type-null", "タイプ：ヌル"},
	{"silvally", "シルヴァディ"},
	{"minior-red-meteor", "メテノ"},
	{"komala", "ネッコアラ"},
	{"turtonator", "バクガメス"},
	{"togedemaru", "トゲデマル"},
	{"mimikyu-disguised", "ミミッキュ"},
	{"bruxish", "ハギギシリ"},
	{"drampa", "ジジーロン"},
	{"dhelmise", "ダダリン"},
	{"jangmo-o", "ジャラコ"},
	{"hakamo-o", "ジャランゴ"},
	{"kommo-o", "ジャラランガ"},
	{"tapu-koko", "カプ・コケコ"},
	{"tapu-lele", "カプ・テテフ"},
	{"tapu-bulu", "カプ・ブルル"},
	{"tapu-fini", "カプ・レヒレ"},
	{"cosmog", "コスモッグ"},
	{"cosmoem", "コスモウム"},
	{"solgaleo", "ソルガレオ"},
	{"lunala", "ルナアーラ"},
	{"nihilego", "ウツロイド"},
	{"buzzwole", "マッシブーン"},
	{"pheromosa", "フェローチェ"},
	{"xurkitree", "デンジュモク"},
	{"celesteela", "テッカグヤ"},
	{"kartana", "カミツルギ"},
	{"guzzlord", "アクジキング"},
	{"necrozma", "ネクロズマ"},
	{"magearna", "マギアナ"},
	{"marshadow", "マーシャドー"},
	{"poipole", "ベベノム"},
	{"naganadel", "アーゴヨン"},
	{"stakataka", "ツンデツンデ"},
	{"blacephalon", "ズガドーン"},
	{"zeraora", "ゼラオラ"},
	{"meltan", "メルタン"},
	{"melmetal", "メルメタル"},
	{"grookey", "サルノリ"},
	{"thwackey", "バチンキー"},
	{"rillaboom", "ゴリランダー"},
	{"scorbunny", "ヒバニー"},
	{"raboot", "ラビフット"},
	{"cinderace", "エースバーン"},
	{"sobble", "メッソン"},
	{"drizzile", "ジメレオン"},
	{"inteleon", "インテレオン"},
	{"skwovet", "ホシガリス"},
	{"greedent", "ヨクバリス"},
	{"rookidee", "ココガラ"},
	{"corvisquire", "アオガラス"},
	{"corviknight", "アーマーガア"},
	{"blipbug", "サッチムシ"},
	{"dottler", "レドームシ"},
	{"orbeetle", "イオルブ"},
	{"nickit", "クスネ"},
	{"thievul", "フォクスライ"},
	{"gossifleur", "ヒメンカ"},
	{"eldegoss", "ワタシラガ"},
	{"wooloo", "ウールー"},
	{"dubwool", "バイウールー"},
	{"chewtle", "カムカメ"},
	{"drednaw", "カジリガメ"},
	{"yamper", "ワンパチ"},
	{"boltund", "パルスワン"},
	{"rolycoly", "タンドン"},
	{"carkol", "トロッゴン"},
	{"coalossal", "セキタンザン"},
	{"applin", "カジッチュ"},
	{"flapple", "アップリュー"},
	{"appletun", "タルップル"},
	{"silicobra", "スナヘビ"},
	{"sandaconda", "サダイジャ"},
	{"cramorant", "ウッウ"},
	{"arrokuda", "サシカマス"},
	{"barraskewda", "カマスジョー"},
	{"toxel", "エレズン"},
	{"toxtricity", "ストリンダー"},
	{"sizzlipede", "ヤクデ"},
	{"centiskorch", "マルヤクデ"},
	{"clobbopus", "タタッコ"},
	{"grapploct", "オトスパス"},
	{"sinistea", "ヤバチャ"},
	{"polteageist", "ポットデス"},
	{"hatenna", "ミブリム"},
	{"hattrem", "テブリム"},
	{"hatterene", "ブリムオン"},
	{"impidimp", "ベロバー"},
	{"morgrem", "ギモー"},
	{"grimmsnarl", "オーロンゲ"},
	{"obstagoon", "タチフサグマ"},
	{"perrserker", "ニャイキング"},
	{"cursola", "サニゴーン"},
	{"sirfetch'd", "ネギガナイト"},
	{"mr. Rime", "バリコオル"},
	{"runerigus", "デスバーン"},
	{"milcery", "マホミル"},
	{"alcremie", "マホイップ"},
	{"falinks", "タイレーツ"},
	{"pincurchin", "バチンウニ"},
	{"snom", "ユキハミ"},
	{"frosmoth", "モスノウ"},
	{"stonjourner", "イシヘンジン"},
	{"eiscue", "コオリッポ"},
	{"indeedee", "イエッサン"},
	{"morpeko", "モルペコ"},
	{"cufant", "ゾウドウ"},
	{"copperajah", "ダイオウドウ"},
	{"dracozolt", "パッチラゴン"},
	{"arctozolt", "パッチルドン"},
	{"dracovish", "ウオノラゴン"},
	{"arctovish", "ウオチルドン"},
	{"duraludon", "ジュラルドン"},
	{"dreepy", "ドラメシヤ"},
	{"drakloak", "ドロンチ"},
	{"dragapult", "ドラパルト"},
	{"zacian", "ザシアン"},
	{"zamazenta", "ザマゼンタ"},
	{"eternatus", "ムゲンダイナ"},
	{"kubfu", "ダクマ"},
	{"urshifu", "ウーラオス"},
	{"zarude", "ザルード"},
	// フォルムチェンジ
	{"rattata-alola", "コラッタ(アローラ)"},
	{"raticate-alola", "ラッタ(アローラ)"},
	{"raichu-alola", "ライチュウ(アローラ)"},
	{"sandshrew-alola", "サンド(アローラ)"},
	{"sandslash-alola", "サンドパン(アローラ)"},
	{"vulpix-alola", "ロコン(アローラ)"},
	{"ninetales-alola", "キュウコン(アローラ)"},
	{"diglett-alola", "ディグダ(アローラ)"},
	{"dugtrio-alola", "ダグトリオ(アローラ)"},
	{"meowth-alola", "ニャース(アローラ)"},
	{"persian-alola", "ペルシアン(アローラ)"},
	{"geodude-alola", "イシツブテ(アローラ)"},
	{"graveler-alola", "ゴローン(アローラ)"},
	{"golem-alola", "ゴローニャ(アローラ)"},
	{"grimer-alola", "ベトベター(アローラ)"},
	{"muk-alola", "ベトベトン(アローラ)"},
	{"exeggutor-alola", "ナッシー(アローラ)"},
	{"marowak-alola", "ガラガラ(アローラ)"},
	{"deoxys-attack", "デオキシス(アタック)"},
	{"deoxys-defense", "デオキシス(ディフェンス)"},
	{"deoxys-speed", "デオキシス(スピード)"},
	{"wormadam-sandy", "ミノマダム(すなちのミノ)"},
	{"wormadam-trash", "ミノマダム(ゴミのミノ)"},
	{"rotom-heat", "ヒートロトム"},
	{"rotom-wash", "ウォッシュロトム"},
	{"rotom-frost", "フロストロトム"},
	{"rotom-fan", "スピンロトム"},
	{"rotom-mow", "カットロトム"},
	{"giratina-origin", "ギラティナオリジン"},
	{"shaymin-sky", "シェイミスカイフォルム"},
	{"basculin-blue-striped", "バスラオ(あおすじ)"},
	{"darmanitan-zen", "ヒヒダルマ(ダルマモード)"},
	{"tornadus-therian", "トルネロス(れいじゅう)"},
	{"thundurus-therian", "ボルトロス(れいじゅう)"},
	{"landorus-therian", "ランドロス(れいじゅう)"},
	{"kyurem-black", "ブラックキュレム"},
	{"kyurem-white", "ホワイトキュレム"},
	{"keldeo-resolute", "ケルディオ(かくごのすがた)"},
	{"meloetta-pirouette", "メロエッタ(ステップフォルム)"},
	{"greninja-ash", "サトシゲッコウガ"},
	{"floette-eternal", "フラエッテ"},
	{"meowstic-female", "ニャオニクス♀"},
	{"aegislash-blade", "ギルガルド(ブレード)"},
	{"pumpkaboo-small", "バケッチャ小"},
	{"pumpkaboo-large", "バケッチャ大"},
	{"pumpkaboo-super", "バケッチャ特大"},
	{"gourgeist-small", "パンプジン小"},
	{"gourgeist-large", "パンプジン大"},
	{"gourgeist-super", "パンプジン特大"},
	{"zygarde-10", "ジガルデ10%"},
	{"zygarde-50", "ジガルデ50%"},
	{"zygarde-complete", "ジガルデ100%"},
	{"hoopa-unbound", "ときはなたれしフーパ"},
	{"oricorio-pom-pom", "オドリドリぱちぱち"},
	{"oricorio-pau", "オドリドリふらふら"},
	{"oricorio-sensu", "オドリドリまいまい"},
	{"lycanroc-midnight", "ルガルガン(まよなか)"},
	{"lycanroc-dusk", "ルガルガン(たそがれ)"},
	{"wishiwashi-school", "ヨワシ(むれ)"},
	{"minior-red", "メテノ(コア)"},
	{"necrozma-dusk", "ネクロズマ(たそがれ)"},
	{"necrozma-dawn", "ネクロズマ(あかつき)"},
	{"necrozma-ultra", "ウルトラネクロズマ"},
}
