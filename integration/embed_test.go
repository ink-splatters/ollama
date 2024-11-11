//go:build integration

package integration

import (
	"context"
	"math"
	"testing"
	"time"

	"github.com/ink-splatters/ollama/api"
)

func dotProduct[V float32 | float64](v1, v2 []V) V {
	var result V = 0
	for i := 0; i < len(v1); i++ {
		result += v1[i] * v2[i]
	}
	return result
}

func magnitude[V float32 | float64](v []V) V {
	var result V = 0
	for _, val := range v {
		result += val * val
	}
	return V(math.Sqrt(float64(result)))
}

func cosineSimilarity[V float32 | float64](v1, v2 []V) V {
	return dotProduct(v1, v2) / (magnitude(v1) * magnitude(v2))
}

func TestAllMiniLMEmbeddings(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	req := api.EmbeddingRequest{
		Model:  "all-minilm",
		Prompt: "why is the sky blue?",
	}

	res, err := embeddingTestHelper(ctx, t, req)

	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if len(res.Embedding) != 384 {
		t.Fatalf("expected 384 floats, got %d", len(res.Embedding))
	}

	expected := []float64{
		0.06642947345972061, -0.01160573959350586, 0.3302811086177826, 0.309552937746048, 0.36223655939102173, 0.05672447010874748, 0.6955016851425171, -0.17069467902183533, 0.8547305464744568, 0.21076075732707977, -0.29339903593063354, -0.05926772207021713, -0.003363408148288727, -0.4204462468624115, -0.1061280220746994, 0.30754348635673523, -0.14551642537117004, -1.0430994033813477, -0.4805174171924591, -0.40448474884033203, -0.4345352053642273, 0.3573606014251709, -0.4098161458969116, 0.25664326548576355, -0.3021087646484375, 0.36236199736595154, -0.23262615501880646, 0.08319848775863647, 0.28042519092559814, -0.052289899438619614, -0.12552005052566528, 0.402255117893219, 0.24357250332832336, 0.08881516754627228, -0.17023836076259613, -0.2868475615978241, 0.4790303707122803, -0.3199635446071625, 0.02826809138059616, -0.19417747855186462, -0.19217649102210999, -0.21705707907676697, -0.1210065633058548, 0.10262420773506165, -0.07726037502288818, 0.10094445943832397, -0.06194962561130524, 0.1712605208158493, 0.628441333770752, -0.10222385078668594, -0.16214007139205933, 0.059920795261859894, -0.5053377151489258, 0.10545563697814941, 0.32686805725097656, 0.7650210857391357, 0.006465774029493332, -0.13403119146823883, 0.6090353727340698, 0.05603303387761116, -0.37635889649391174, 0.45424884557724, -0.5053073763847351, 0.4572359323501587, 0.6084011197090149, -0.3659921884536743, -0.3536888360977173, 0.05569244921207428, -0.4166066646575928, -0.43796032667160034, -0.16600576043128967, 0.12460685521364212, 0.40493422746658325, -0.18632565438747406, 0.2390710711479187, 0.007283639162778854, 0.4001992344856262, -0.4455743134021759, -0.05360018089413643, -0.08401738107204437, 0.2041706144809723, -0.42083415389060974, -0.491476833820343, 0.7860275506973267, 0.08280622214078903, 0.4309011697769165, 0.09778489172458649, 0.3392091989517212, -0.5618907809257507, 0.06766007840633392, -0.05127308890223503, -0.23472431302070618, -0.7611223459243774, -0.20227840542793274, -0.5491426587104797, 0.09030043333768845, 0.37326449155807495, -0.2696656584739685, 0.2814738154411316, 0.1461343765258789, 0.309052437543869, -0.3387487828731537, 0.1990429162979126, 0.0474909171462059, -0.02756538614630699, -0.20544570684432983, 0.5137258768081665, 0.22562497854232788, 0.40487033128738403, 0.04954294115304947, -0.23911823332309723, -0.5578761696815491, 0.14376327395439148, -0.12795016169548035, -0.26285219192504883, 0.3614377975463867, -0.22225692868232727, 0.11940789222717285, -0.6961514353752136, -0.3324243426322937, -0.07613810151815414, 0.24946099519729614, 0.1462409496307373, 0.5309336185455322, 0.051560595631599426, -0.11104149371385574, -0.39189594984054565, -4.767201176712463e-32, 0.892546534538269, -0.07396792620420456, 0.6088366508483887, 0.23729179799556732, 0.2614588737487793, -0.3626874089241028, -0.23131835460662842, -0.024579279124736786, -0.12901946902275085, -0.2306443750858307, -0.0376533679664135, -0.09649471938610077, -0.16013199090957642, -0.31914401054382324, 0.3151017129421234, -0.11264121532440186, -0.4020160734653473, 0.039211247116327286, -0.5478582978248596, 0.5563258528709412, -0.6903842091560364, 0.2746567130088806, -0.24196553230285645, -0.053318753838539124, -0.18611761927604675, -0.28490889072418213, 0.237456813454628, 0.4946249723434448, 0.37237465381622314, 0.07815749943256378, 0.6494859457015991, 0.6915512084960938, -0.14422327280044556, 0.30338582396507263, -0.17378094792366028, -0.33589833974838257, -0.09702004492282867, -0.04210608825087547, -0.566387414932251, 0.18866634368896484, -0.3533778488636017, 0.37286972999572754, -0.39420801401138306, 0.0818595215678215, 0.436712384223938, -0.08886678516864777, 0.2527940273284912, -0.5864061117172241, -0.37891554832458496, 0.21103361248970032, -0.2275354266166687, 0.1558678150177002, 0.09536703675985336, -0.27437490224838257, 0.4484926164150238, 0.20584626495838165, 0.45972558856010437, -0.231113001704216, -0.021833699196577072, 0.3253912925720215, -0.08802174031734467, -0.023067735135555267, 0.33492740988731384, 0.5189340114593506, 0.2481488585472107, -0.07638847082853317, 0.25147074460983276, 0.2771286964416504, -0.08443005383014679, -0.5207436084747314, 0.05951530486345291, 0.08816319704055786, 0.15935833752155304, 0.0644921213388443, -0.07194079458713531, -0.5383226871490479, 0.17800968885421753, -0.195652037858963, -0.028597159311175346, 0.08582349121570587, -0.23225288093090057, -0.12984338402748108, 0.3651025593280792, -0.4039592146873474, -0.3628298342227936, 0.08263863623142242, -0.12648534774780273, -0.08284908533096313, -0.1042669266462326, -0.4579034447669983, -0.2961195111274719, -0.32282471656799316, 0.3182551860809326, -0.6890494227409363, -0.7114676237106323, 2.3665072841905432e-32, -0.0030965525656938553, -0.5696439146995544, -0.5794872045516968, 0.04729880392551422, -0.048917483538389206, -0.10963250696659088, 0.298623263835907, 0.4452674388885498, -0.2828809320926666, 0.5696343183517456, 0.3004711866378784, 0.44842660427093506, 0.06550214439630508, -0.020054858177900314, 0.385932058095932, -0.23460465669631958, 0.23865005373954773, 0.4363722801208496, -0.24931970238685608, -0.41073542833328247, -0.2937365770339966, 0.5095447301864624, 0.2864843010902405, -0.14028388261795044, -0.14269764721393585, 0.4107881486415863, -0.2581801116466522, 0.18544888496398926, -0.08612997084856033, 0.33715111017227173, -0.24288496375083923, 0.3599962592124939, -0.43829354643821716, 0.15094976127147675, 0.03177203983068466, 0.5965112447738647, 0.03364168107509613, -0.5481097102165222, -0.363423228263855, 0.4825053811073303, -0.7288467288017273, -0.13361915946006775, 0.7423286437988281, -0.3515661358833313, -0.37989044189453125, -0.1576842963695526, 0.3734908998012543, 0.8393698930740356, 0.23719121515750885, -0.28990280628204346, 0.11215505003929138, -0.16382968425750732, 0.47951722145080566, 0.28471529483795166, 0.5308315753936768, -0.1286555975675583, -0.22689077258110046, 0.6377706527709961, 0.34224453568458557, 0.07091143727302551, 0.26538553833961487, 0.014475930482149124, -0.050034329295158386, 0.011025313287973404, 0.09357182681560516, 0.1345357596874237, -0.1523902863264084, 0.14176052808761597, -0.0609259307384491, -0.3332745134830475, -0.1072426363825798, -0.5933747291564941, -0.40028926730155945, 0.5343422293663025, 0.016202416270971298, 0.27436596155166626, 0.28844428062438965, -0.1660136878490448, -0.6286065578460693, 0.5850632190704346, -0.6491153836250305, -0.03207448124885559, 0.23312292993068695, 0.09339666366577148, -0.42595869302749634, -0.5011518001556396, 0.08187201619148254, -0.3312609791755676, -0.3677852153778076, -0.3758619427680969, -0.12195874005556107, -0.014479270204901695, -0.014539752155542374, 0.23270025849342346, -0.3609132170677185, -9.438503667524856e-8, -0.05230816453695297, 0.17612962424755096, 0.01489749364554882, 0.06601762771606445, -0.14300350844860077, -0.1422577053308487, 0.7347333431243896, 0.030603498220443726, 0.24959787726402283, 0.026135217398405075, -0.4412609338760376, -0.18663707375526428, -0.29235413670539856, 0.4696626365184784, 0.12353914976119995, -0.3236965537071228, -0.6856554746627808, -0.28768694400787354, 0.0671629011631012, 0.27566438913345337, -0.0893339067697525, -0.22328855097293854, -0.16536207497119904, -0.08968719840049744, 0.022607458755373955, 0.21818216145038605, -0.14408129453659058, 0.14458191394805908, 0.4712568521499634, 0.13527995347976685, 0.16118602454662323, 0.23675017058849335, -0.0062652211636304855, -0.4045848250389099, -0.5631943345069885, 0.04897312819957733, -0.2558498978614807, 0.5269845128059387, -0.16870160400867462, -0.39874112606048584, 0.3996037244796753, 0.5432316660881042, -0.3740345239639282, 0.031965695321559906, 0.29769593477249146, 0.1568443477153778, 0.287019282579422, 0.6005253791809082, -0.33905476331710815, -0.07407552748918533, -0.4541633129119873, 0.047827333211898804, 0.4803982973098755, -0.2860602140426636, 0.17097190022468567, -0.7525586485862732, -0.06290972977876663, 0.14645379781723022, 0.176426962018013, 0.024587953463196754, 0.105128213763237, 0.023733407258987427, -0.1363760083913803, 0.22127331793308258,
	}
	sim := cosineSimilarity(res.Embedding, expected)
	if sim < 0.99 {
		t.Fatalf("expected %v, got %v (similarity: %f)", expected[0:5], res.Embedding[0:5], sim)
	}
}

func TestAllMiniLMEmbed(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	req := api.EmbedRequest{
		Model: "all-minilm",
		Input: "why is the sky blue?",
	}

	res, err := embedTestHelper(ctx, t, req)

	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if len(res.Embeddings) != 1 {
		t.Fatalf("expected 1 embedding, got %d", len(res.Embeddings))
	}

	if len(res.Embeddings[0]) != 384 {
		t.Fatalf("expected 384 floats, got %d", len(res.Embeddings[0]))
	}

	expected := []float32{
		0.010071031, -0.0017594865, 0.050072223, 0.046929732, 0.05491682, 0.008599705, 0.105441436, -0.025878143, 0.1295813, 0.031952355, -0.04448072, -0.0089852745, -0.000509909, -0.06374169, -0.016089523, 0.04662509, -0.022060998, -0.15813895, -0.072848774, -0.061321855, -0.065877646, 0.054177605, -0.06213012, 0.038908366, -0.04580116, 0.05493584, -0.035267256, 0.012613296, 0.04251382, -0.007927403, -0.01902945, 0.060983833, 0.036926776, 0.013464811, -0.025808964, -0.043487485, 0.072623335, -0.04850803, 0.00428558, -0.02943825, -0.02913489, -0.03290691, -0.018345183, 0.0155583285, -0.011713048, 0.01530367, -0.009391865, 0.025963927, 0.09527476, -0.015497632, -0.024581224, 0.009084283, -0.07661165, 0.015987588, 0.049554788, 0.115980916, 0.0009802427, -0.02031978, 0.09233272, 0.00849488, -0.05705784, 0.068866335, -0.076607056, 0.06931919, 0.09223656, -0.055486195, -0.053620946, 0.008443246, -0.06315959, -0.066396914, -0.02516728, 0.018891005, 0.061389998, -0.028247874, 0.036244337, 0.0011042351, 0.06067215, -0.06755123, -0.008126048, -0.012737444, 0.030953258, -0.06380051, -0.07451028, 0.1191656, 0.012553826, 0.06532671, 0.014824665, 0.051425762, -0.08518537, 0.010257597, -0.0077732494, -0.035585348, -0.115389846, -0.03066639, -0.0832527, 0.013689985, 0.056588713, -0.040882625, 0.042672798, 0.022154681, 0.04685385, -0.05135596, 0.030175874, 0.007199854, -0.0041790465, -0.031146567, 0.07788334, 0.034205843, 0.06138031, 0.007510951, -0.036251485, -0.08457674, 0.021795211, -0.019397866, -0.03984967, 0.054795727, -0.033695232, 0.018102817, -0.10553994, -0.050397146, -0.011542906, 0.0378195, 0.022170838, 0.08049212, 0.007816837, -0.01683443, -0.059413332, -7.227309e-33, 0.13531439, -0.011213897, 0.0923026, 0.03597459, 0.039638437, -0.054985173, -0.03506899, -0.0037263383, -0.01955998, -0.034966808, -0.0057084337, -0.014629069, -0.024276787, -0.048383784, 0.04777095, -0.017076956, -0.06094759, 0.0059446157, -0.083057985, 0.084341705, -0.1046656, 0.041639294, -0.03668315, -0.008083383, -0.028216336, -0.04319357, 0.035999607, 0.07498755, 0.05645381, 0.011849057, 0.09846523, 0.10484252, -0.021864949, 0.045994766, -0.026346037, -0.05092382, -0.014708711, -0.0063834875, -0.085867085, 0.028602734, -0.0535738, 0.056528863, -0.059763853, 0.012410302, 0.06620772, -0.013472636, 0.038324803, -0.08890202, -0.05744544, 0.03199372, -0.034495477, 0.02363032, 0.014458106, -0.04159657, 0.06799366, 0.031207295, 0.069696635, -0.035037853, -0.0033100948, 0.0493309, -0.0133445235, -0.0034971808, 0.050776623, 0.078672916, 0.037620574, -0.011580864, 0.03812419, 0.04201406, -0.012800006, -0.07894726, 0.00902281, 0.013365969, 0.024159499, 0.009777319, -0.010906574, -0.08161233, 0.026987134, -0.0296618, -0.004335468, 0.013011258, -0.035210665, -0.019684888, 0.055351324, -0.06124218, -0.055006765, 0.012528419, -0.019175794, -0.012560324, -0.015807373, -0.06942039, -0.044893157, -0.048941795, 0.048249032, -0.10446324, -0.10786195, 3.58774e-33, -0.0004694524, -0.08636079, -0.087853074, 0.0071707284, -0.007416128, -0.01662082, 0.045272738, 0.06750471, -0.042886123, 0.08635933, 0.04555289, 0.06798365, 0.009930444, -0.003040414, 0.058509175, -0.035567205, 0.036180507, 0.06615616, -0.03779808, -0.062269486, -0.044531893, 0.07724946, 0.04343241, -0.021267718, -0.021633657, 0.06227748, -0.03914136, 0.028114952, -0.013057723, 0.051113747, -0.036822543, 0.054577183, -0.06644743, 0.022884717, 0.0048167957, 0.09043401, 0.0051002423, -0.083096094, -0.055096727, 0.07315016, -0.11049671, -0.020257315, 0.11254063, -0.053299136, -0.057593238, -0.023905706, 0.056623034, 0.12725255, 0.03595934, -0.043950673, 0.017003251, -0.024837377, 0.07269714, 0.043164223, 0.08047665, -0.019504813, -0.034397744, 0.096689135, 0.051885936, 0.010750518, 0.04023374, 0.0021946214, -0.0075854477, 0.0016714911, 0.014185944, 0.020396275, -0.023103109, 0.021491585, -0.009236667, -0.050526038, -0.016258504, -0.0899585, -0.0606858, 0.08100888, 0.0024563652, 0.041595213, 0.043729555, -0.025168482, -0.09529981, 0.088698424, -0.09840905, -0.0048626475, 0.03534257, 0.014159388, -0.06457741, -0.07597705, 0.012412196, -0.050220776, -0.055758025, -0.0569825, -0.018489538, -0.0021951278, -0.002204297, 0.03527849, -0.0547162, -1.430923e-8, -0.007930172, 0.026702108, 0.0022585324, 0.010008593, -0.021680027, -0.02156696, 0.111389145, 0.004639639, 0.03784025, 0.003962226, -0.0668973, -0.028295087, -0.04432231, 0.07120314, 0.018729135, -0.04907397, -0.103948705, -0.043614738, 0.010182222, 0.04179206, -0.013543455, -0.03385163, -0.025069695, -0.013597015, 0.0034274007, 0.033077475, -0.021843424, 0.021919321, 0.07144483, 0.020509098, 0.024436586, 0.035892475, -0.00094983797, -0.061337028, -0.085383, 0.007424564, -0.038788088, 0.07989341, -0.025575982, -0.060451094, 0.060581867, 0.082356565, -0.056705453, 0.0048461547, 0.04513215, 0.023778366, 0.043513518, 0.09104256, -0.05140235, -0.01123021, -0.06885336, 0.007250856, 0.072830714, -0.04336812, 0.025920171, -0.11409155, -0.009537421, 0.022203108, 0.026747186, 0.0037276533, 0.015937949, 0.0035980998, -0.020675266, 0.03354611,
	}
	sim := cosineSimilarity(res.Embeddings[0], expected)
	if sim < 0.99 {
		t.Fatalf("expected %v, got %v (similarity: %f)", expected[0:5], res.Embeddings[0][0:5], sim)
	}

	if res.PromptEvalCount != 6 {
		t.Fatalf("expected 6 prompt tokens, got %d", res.PromptEvalCount)
	}
}

func TestAllMiniLMBatchEmbed(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	req := api.EmbedRequest{
		Model: "all-minilm",
		Input: []string{"why is the sky blue?", "why is the grass green?"},
	}

	res, err := embedTestHelper(ctx, t, req)

	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if len(res.Embeddings) != 2 {
		t.Fatalf("expected 2 embeddings, got %d", len(res.Embeddings))
	}

	if len(res.Embeddings[0]) != 384 {
		t.Fatalf("expected 384 floats, got %d", len(res.Embeddings[0]))
	}

	expected := [][]float32{
		{
			0.010071031, -0.0017594865, 0.050072223, 0.046929732, 0.05491682, 0.008599705, 0.105441436, -0.025878143, 0.1295813, 0.031952355, -0.04448072, -0.0089852745, -0.000509909, -0.06374169, -0.016089523, 0.04662509, -0.022060998, -0.15813895, -0.072848774, -0.061321855, -0.065877646, 0.054177605, -0.06213012, 0.038908366, -0.04580116, 0.05493584, -0.035267256, 0.012613296, 0.04251382, -0.007927403, -0.01902945, 0.060983833, 0.036926776, 0.013464811, -0.025808964, -0.043487485, 0.072623335, -0.04850803, 0.00428558, -0.02943825, -0.02913489, -0.03290691, -0.018345183, 0.0155583285, -0.011713048, 0.01530367, -0.009391865, 0.025963927, 0.09527476, -0.015497632, -0.024581224, 0.009084283, -0.07661165, 0.015987588, 0.049554788, 0.115980916, 0.0009802427, -0.02031978, 0.09233272, 0.00849488, -0.05705784, 0.068866335, -0.076607056, 0.06931919, 0.09223656, -0.055486195, -0.053620946, 0.008443246, -0.06315959, -0.066396914, -0.02516728, 0.018891005, 0.061389998, -0.028247874, 0.036244337, 0.0011042351, 0.06067215, -0.06755123, -0.008126048, -0.012737444, 0.030953258, -0.06380051, -0.07451028, 0.1191656, 0.012553826, 0.06532671, 0.014824665, 0.051425762, -0.08518537, 0.010257597, -0.0077732494, -0.035585348, -0.115389846, -0.03066639, -0.0832527, 0.013689985, 0.056588713, -0.040882625, 0.042672798, 0.022154681, 0.04685385, -0.05135596, 0.030175874, 0.007199854, -0.0041790465, -0.031146567, 0.07788334, 0.034205843, 0.06138031, 0.007510951, -0.036251485, -0.08457674, 0.021795211, -0.019397866, -0.03984967, 0.054795727, -0.033695232, 0.018102817, -0.10553994, -0.050397146, -0.011542906, 0.0378195, 0.022170838, 0.08049212, 0.007816837, -0.01683443, -0.059413332, -7.227309e-33, 0.13531439, -0.011213897, 0.0923026, 0.03597459, 0.039638437, -0.054985173, -0.03506899, -0.0037263383, -0.01955998, -0.034966808, -0.0057084337, -0.014629069, -0.024276787, -0.048383784, 0.04777095, -0.017076956, -0.06094759, 0.0059446157, -0.083057985, 0.084341705, -0.1046656, 0.041639294, -0.03668315, -0.008083383, -0.028216336, -0.04319357, 0.035999607, 0.07498755, 0.05645381, 0.011849057, 0.09846523, 0.10484252, -0.021864949, 0.045994766, -0.026346037, -0.05092382, -0.014708711, -0.0063834875, -0.085867085, 0.028602734, -0.0535738, 0.056528863, -0.059763853, 0.012410302, 0.06620772, -0.013472636, 0.038324803, -0.08890202, -0.05744544, 0.03199372, -0.034495477, 0.02363032, 0.014458106, -0.04159657, 0.06799366, 0.031207295, 0.069696635, -0.035037853, -0.0033100948, 0.0493309, -0.0133445235, -0.0034971808, 0.050776623, 0.078672916, 0.037620574, -0.011580864, 0.03812419, 0.04201406, -0.012800006, -0.07894726, 0.00902281, 0.013365969, 0.024159499, 0.009777319, -0.010906574, -0.08161233, 0.026987134, -0.0296618, -0.004335468, 0.013011258, -0.035210665, -0.019684888, 0.055351324, -0.06124218, -0.055006765, 0.012528419, -0.019175794, -0.012560324, -0.015807373, -0.06942039, -0.044893157, -0.048941795, 0.048249032, -0.10446324, -0.10786195, 3.58774e-33, -0.0004694524, -0.08636079, -0.087853074, 0.0071707284, -0.007416128, -0.01662082, 0.045272738, 0.06750471, -0.042886123, 0.08635933, 0.04555289, 0.06798365, 0.009930444, -0.003040414, 0.058509175, -0.035567205, 0.036180507, 0.06615616, -0.03779808, -0.062269486, -0.044531893, 0.07724946, 0.04343241, -0.021267718, -0.021633657, 0.06227748, -0.03914136, 0.028114952, -0.013057723, 0.051113747, -0.036822543, 0.054577183, -0.06644743, 0.022884717, 0.0048167957, 0.09043401, 0.0051002423, -0.083096094, -0.055096727, 0.07315016, -0.11049671, -0.020257315, 0.11254063, -0.053299136, -0.057593238, -0.023905706, 0.056623034, 0.12725255, 0.03595934, -0.043950673, 0.017003251, -0.024837377, 0.07269714, 0.043164223, 0.08047665, -0.019504813, -0.034397744, 0.096689135, 0.051885936, 0.010750518, 0.04023374, 0.0021946214, -0.0075854477, 0.0016714911, 0.014185944, 0.020396275, -0.023103109, 0.021491585, -0.009236667, -0.050526038, -0.016258504, -0.0899585, -0.0606858, 0.08100888, 0.0024563652, 0.041595213, 0.043729555, -0.025168482, -0.09529981, 0.088698424, -0.09840905, -0.0048626475, 0.03534257, 0.014159388, -0.06457741, -0.07597705, 0.012412196, -0.050220776, -0.055758025, -0.0569825, -0.018489538, -0.0021951278, -0.002204297, 0.03527849, -0.0547162, -1.430923e-8, -0.007930172, 0.026702108, 0.0022585324, 0.010008593, -0.021680027, -0.02156696, 0.111389145, 0.004639639, 0.03784025, 0.003962226, -0.0668973, -0.028295087, -0.04432231, 0.07120314, 0.018729135, -0.04907397, -0.103948705, -0.043614738, 0.010182222, 0.04179206, -0.013543455, -0.03385163, -0.025069695, -0.013597015, 0.0034274007, 0.033077475, -0.021843424, 0.021919321, 0.07144483, 0.020509098, 0.024436586, 0.035892475, -0.00094983797, -0.061337028, -0.085383, 0.007424564, -0.038788088, 0.07989341, -0.025575982, -0.060451094, 0.060581867, 0.082356565, -0.056705453, 0.0048461547, 0.04513215, 0.023778366, 0.043513518, 0.09104256, -0.05140235, -0.01123021, -0.06885336, 0.007250856, 0.072830714, -0.04336812, 0.025920171, -0.11409155, -0.009537421, 0.022203108, 0.026747186, 0.0037276533, 0.015937949, 0.0035980998, -0.020675266, 0.03354611,
		},
		{
			-0.009802706, 0.060424678, 0.025257956, -0.0063643856, 0.07272723, 0.01719488, 0.090320334, -0.051705167, 0.099515095, 0.09072479, 0.007301506, -0.01968127, -0.075095184, -0.017409375, 0.019365614, 0.040805466, -0.011079843, -0.05856395, -0.12545314, -0.048980292, -0.044052314, 0.03115607, 0.037880868, -0.03187379, -0.0909825, 0.06357952, -0.076541565, 0.085011445, 0.03554875, -0.071272224, 0.021114277, 0.11005397, 0.03312636, -0.025947863, -0.061563145, -0.026466936, 0.02054478, -0.05426622, 0.056569945, 0.03292456, -0.09005933, -0.05698778, 0.026827272, 0.0751872, -0.07142025, -0.0043633, 0.054151993, 0.026441583, 0.078053534, -0.048995998, 0.056577347, -0.048973206, -0.07581186, 0.006902122, 0.0062451144, 0.037024222, 0.025028007, 0.021724675, 0.010117283, -0.040492155, -0.012010403, -0.03334674, -0.07570402, 0.071321115, -0.02062346, -0.0631419, -0.001237942, -0.055173304, 0.009124682, -0.08703634, 0.020684991, 0.05294139, -0.009563882, -0.052647192, -0.06467313, 0.041968923, 0.04473555, 0.03270584, -0.019611169, 0.00013324046, 0.038228948, 0.0509972, 0.0047100335, 0.05736671, 0.046469305, 0.04269017, -0.017305125, 0.011859765, -0.05701112, -0.03498464, -0.018940303, -0.0074608736, -0.07385685, 0.043892473, -0.09890047, 0.041379265, -0.024019944, -0.12034819, 0.0001821356, -0.0038607453, 0.056144036, -0.0005059898, 0.07110965, -0.03616245, -0.06406574, -0.009435536, -0.042290587, 0.07791005, -0.02365763, 0.007864432, -0.023739463, -0.018536761, -0.033538047, 0.0776669, -0.06058719, 0.05363198, 0.033863083, 0.012545284, -0.03260245, 0.029770961, -0.016934512, 0.028213669, -0.018053731, 0.06651968, -0.06952628, -0.017853932, -0.037421644, -6.839719e-33, -0.0055490523, -0.031681225, 0.04819487, -0.09944883, 0.09372583, -0.051811725, -0.037059266, -0.026262678, -0.037466466, -0.030253021, 0.0060922937, -0.09831781, -0.017570594, -0.07247917, 0.03856134, 0.00888377, -0.13072893, 0.02145255, -0.075681135, -0.010470858, -0.017236665, 0.058358245, 0.022016024, 0.0015762328, 0.009419801, -0.031423207, 0.08002972, 0.030580623, 0.05696977, -0.012164853, 0.11575935, 0.0040441174, 0.01759827, 0.043209996, 0.02948431, -0.0069428794, -0.025078153, -0.026160793, 0.013364178, 0.121543564, -0.004469769, -0.04534167, 0.043418996, -0.01768049, 0.062162045, -0.039375506, 0.017406953, 0.008458191, -0.02603069, 0.010130821, 0.023227274, 0.05305319, 0.06899141, 0.053088874, -0.0003113895, 0.009642751, 0.08884011, -0.030399954, -0.090916164, -0.051467095, -0.07382789, 0.08624027, 0.003223033, 0.010827092, -0.008318035, -0.011421701, -0.02900046, 0.06548931, 0.005405483, 0.068780296, 0.0428464, -0.01878741, -0.016996592, -0.036818627, -0.0062817424, -0.08700542, -0.008640271, -0.013171244, -0.004574588, 0.04233393, -0.03579696, 0.017357353, -0.087162524, -0.050884914, -0.14957926, -0.002008126, -0.02634847, 0.018098367, 0.02162604, -0.01503002, 0.0037868456, -0.015445877, -0.013303974, -0.09810386, -0.011673153, 2.8261164e-33, -0.022961555, 0.0090464745, -0.0057421196, 0.06604244, 0.042683356, -0.039691485, 0.027226122, 0.03183442, -0.028517157, 0.045575514, -0.055865873, 0.0924774, -0.046869125, 0.08027759, 0.118624836, 0.04889292, -0.06734586, 0.10688813, 0.009396721, -0.051344905, -0.067946814, 0.01592692, -0.010147019, 0.044173665, -0.030018767, 0.022772646, -0.031494025, -0.02233876, -0.0023573847, -0.010024354, 0.0032828946, -0.036839407, -0.11200184, 0.028629173, 0.030212566, 0.03185506, -0.01746865, -0.018295743, -0.036361173, 0.083925165, 0.007943152, -0.023664381, 0.15850149, 0.032088134, -0.070371404, -0.034124147, -0.015502377, 0.07960292, -0.06218589, 0.046537183, 0.04505064, 0.1043822, 0.029607052, 0.047920443, 0.09711685, -0.015767856, -0.064267434, 0.01960162, -0.093837254, -0.0028061024, 0.019721054, -0.027095793, -0.078636706, 0.0689579, 0.107794516, -0.033122607, -0.064406104, 0.016571952, 0.019280795, -0.023045482, -0.018821374, -0.018646069, -0.06431513, -0.03231013, -0.0027636476, 0.059007723, 0.059882853, -0.044795096, -0.06667144, 0.043793377, -0.019855661, -0.006715758, 0.04733659, -0.046866804, 0.03461545, -0.015199261, -0.039511763, 0.047361404, 0.052113988, 0.0008203065, 0.05290727, 0.02459614, -0.029357709, 0.034541644, 0.013009169, -1.36748e-8, -0.033930536, 0.007378359, -0.010701883, 0.04323486, 0.014735074, -0.04162692, 0.10553509, -0.012822099, -0.002357336, 0.040418625, -0.08136588, 0.033679843, -0.019665385, 0.077529214, 0.060347307, -0.016181026, -0.11332622, -0.04306442, 0.023209568, 0.07448782, -0.06055759, -0.045812756, -0.087526724, 0.0534105, -0.044014834, 0.029827949, 0.038628686, 0.016933717, 0.027725562, 0.078133695, 0.055581007, 0.05306717, -0.010792625, -0.029803185, -0.08492531, -0.016416015, 0.030501937, 0.06944753, -0.061944496, -0.122021444, 0.011901371, 0.07258673, -0.017778289, 0.0030972173, 0.014411535, -0.03802866, -0.052976213, 0.060414705, -0.053164586, 0.01794129, -0.104411006, 0.010633235, 0.042881854, 0.042603284, -0.003009017, -0.08530093, -0.039561126, -0.004481811, 0.013104284, -0.008498699, -0.028943708, -0.03587923, 0.05940551, -0.000055299755,
		},
	}

	sim := cosineSimilarity(res.Embeddings[0], expected[0])
	if sim < 0.99 {
		t.Fatalf("expected %v, got %v (similarity: %f)", expected[0][0:5], res.Embeddings[0][0:5], sim)
	}
	sim = cosineSimilarity(res.Embeddings[1], expected[1])
	if sim < 0.99 {
		t.Fatalf("expected %v, got %v (similarity: %f)", expected[1][0:5], res.Embeddings[1][0:5], sim)
	}

	if res.PromptEvalCount != 12 {
		t.Fatalf("expected 12 prompt tokens, got %d", res.PromptEvalCount)
	}
}

func TestAllMiniLMEmbedTruncate(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	truncTrue, truncFalse := true, false

	type testReq struct {
		Name    string
		Request api.EmbedRequest
	}

	reqs := []testReq{
		{
			Name: "Target Truncation",
			Request: api.EmbedRequest{
				Model: "all-minilm",
				Input: "why",
			},
		},
		{
			Name: "Default Truncate",
			Request: api.EmbedRequest{
				Model:   "all-minilm",
				Input:   "why is the sky blue?",
				Options: map[string]any{"num_ctx": 1},
			},
		},
		{
			Name: "Explicit Truncate",
			Request: api.EmbedRequest{
				Model:    "all-minilm",
				Input:    "why is the sky blue?",
				Truncate: &truncTrue,
				Options:  map[string]any{"num_ctx": 1},
			},
		},
	}

	res := make(map[string]*api.EmbedResponse)

	for _, req := range reqs {
		response, err := embedTestHelper(ctx, t, req.Request)
		if err != nil {
			t.Fatalf("error: %v", err)
		}
		res[req.Name] = response
	}

	if res["Target Truncation"].Embeddings[0][0] != res["Default Truncate"].Embeddings[0][0] {
		t.Fatal("expected default request to truncate correctly")
	}

	if res["Default Truncate"].Embeddings[0][0] != res["Explicit Truncate"].Embeddings[0][0] {
		t.Fatal("expected default request and truncate true request to be the same")
	}

	// check that truncate set to false returns an error if context length is exceeded
	_, err := embedTestHelper(ctx, t, api.EmbedRequest{
		Model:    "all-minilm",
		Input:    "why is the sky blue?",
		Truncate: &truncFalse,
		Options:  map[string]any{"num_ctx": 1},
	})

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func embeddingTestHelper(ctx context.Context, t *testing.T, req api.EmbeddingRequest) (*api.EmbeddingResponse, error) {
	client, _, cleanup := InitServerConnection(ctx, t)
	defer cleanup()
	if err := PullIfMissing(ctx, client, req.Model); err != nil {
		t.Fatalf("failed to pull model %s: %v", req.Model, err)
	}

	response, err := client.Embeddings(ctx, &req)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func embedTestHelper(ctx context.Context, t *testing.T, req api.EmbedRequest) (*api.EmbedResponse, error) {
	client, _, cleanup := InitServerConnection(ctx, t)
	defer cleanup()
	if err := PullIfMissing(ctx, client, req.Model); err != nil {
		t.Fatalf("failed to pull model %s: %v", req.Model, err)
	}

	response, err := client.Embed(ctx, &req)

	if err != nil {
		return nil, err
	}

	return response, nil
}
