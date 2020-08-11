package adstxt

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"golang.org/x/net/publicsuffix"
)

// AdSystems holds a list of all known ad systems (i.e. SSPs/exchanges). There is no order or meaning implied by the ID,
// it is merely an auto incrementing number. "CANONICAL_DOMAIN" is the domain that the exchange has declared to be canonical
// (i.e. what should be used in ads.txt files). Where it is not known, CANONICAL_DOMAIN is NULL.
var adSystems map[int]*adSystem

// AdSystemDomains holds  list of all known domains from field #1 of publisher ads.txt files.
// "ID" is a foreign key referencing the ad system's ID from the "adsystem" table
var adSystemDomains map[string]*adSystemDomain

// normalizeMaappingURL holds list of Ads.txt known advertising systems
const normalizeMaappingURL = "https://wiki.iabtechlab.com/index.php?title=Ads.txt_Normalization_Mappings"

func init() {
	adSystems = map[int]*adSystem{
		1:   newAdSystem(1, "Rubicon Project", "rubiconproject.com"),
		2:   newAdSystem(2, "33Across", ""),
		3:   newAdSystem(3, "PubMatic", "pubmatic.com"),
		4:   newAdSystem(4, "OpenX", "openx.com"),
		5:   newAdSystem(5, "Facebook", ""),
		6:   newAdSystem(6, "GumGum", ""),
		7:   newAdSystem(7, "Kargo", ""),
		8:   newAdSystem(8, "Google", "google.com"),
		9:   newAdSystem(9, "bRealtime", ""),
		10:  newAdSystem(10, "Amazon", ""),
		11:  newAdSystem(11, "One by AOL: Display", "adtech.com, aolcloud.net"),
		12:  newAdSystem(12, "LiveIntent", ""),
		13:  newAdSystem(13, "Yieldmo", ""),
		14:  newAdSystem(14, "MoPub", ""),
		15:  newAdSystem(15, "One by AOL: Mobile", "aol.com"),
		16:  newAdSystem(16, "SmartStream", ""),
		17:  newAdSystem(17, "Smaato", ""),
		18:  newAdSystem(18, "Taboola", ""),
		19:  newAdSystem(19, "TrustX", ""),
		20:  newAdSystem(20, "LKQD", ""),
		21:  newAdSystem(21, "Criteo", ""),
		22:  newAdSystem(22, "Exponential", ""),
		23:  newAdSystem(23, "Sovrn", ""),
		24:  newAdSystem(24, "RhythmOne", ""),
		25:  newAdSystem(25, "Yieldbot", ""),
		26:  newAdSystem(26, "Technorati", ""),
		27:  newAdSystem(27, "Bidfluence", ""),
		28:  newAdSystem(28, "Switch Concepts", ""),
		29:  newAdSystem(29, "BrightRoll from Yahoo!", "btrll.com"),
		30:  newAdSystem(30, "Conversant", ""),
		31:  newAdSystem(31, "Sonobi", ""),
		32:  newAdSystem(32, "Spoutable", ""),
		33:  newAdSystem(33, "FreeWheel", "freewheel.tv"),
		34:  newAdSystem(34, "Connatix", ""),
		35:  newAdSystem(35, "Centro Brand Exchange", ""),
		36:  newAdSystem(36, "Positive Mobile", ""),
		37:  newAdSystem(37, "MemeGlobal", ""),
		38:  newAdSystem(38, "Kixer", ""),
		39:  newAdSystem(39, "Sekindo", ""),
		40:  newAdSystem(40, "Improve Digital", "improvedigital.com"),
		41:  newAdSystem(41, "AdForm", ""),
		42:  newAdSystem(42, "MADS", ""),
		43:  newAdSystem(43, "Inneractive", "inner-active.com"),
		44:  newAdSystem(44, "SpotX", "spotx.tv,spotxchange.com"),
		45:  newAdSystem(45, "StreamRail", ""),
		46:  newAdSystem(46, "MediaMath", ""),
		47:  newAdSystem(47, "AdYouLike", ""),
		48:  newAdSystem(48, "Index Exchange", "indexexchange.com"),
		49:  newAdSystem(49, "e-Planning", ""),
		50:  newAdSystem(50, "Kiosked", ""),
		51:  newAdSystem(51, "UnrulyX", ""),
		52:  newAdSystem(52, "Brightcom", ""),
		53:  newAdSystem(53, "PowerInbox", ""),
		54:  newAdSystem(54, "Fyber", "fyber.com"),
		55:  newAdSystem(55, "TidalTV", ""),
		56:  newAdSystem(56, "Nativo", ""),
		57:  newAdSystem(57, "Media.net", ""),
		58:  newAdSystem(58, "YuMe", ""),
		59:  newAdSystem(59, "RevContent", ""),
		60:  newAdSystem(60, "Outbrain", ""),
		61:  newAdSystem(61, "Zedo", "zedo.com"),
		62:  newAdSystem(62, "SlimCut Media", ""),
		63:  newAdSystem(63, "Bidtellect", ""),
		64:  newAdSystem(64, "Smart RTB+", "smartadserver.com"),
		65:  newAdSystem(65, "LoopMe", "loopme.com"),
		66:  newAdSystem(66, "Vidazoo", ""),
		67:  newAdSystem(67, "Videoflare", ""),
		68:  newAdSystem(68, "Gemini from Yahoo!", "yahoo.com"),
		69:  newAdSystem(69, "PixFuture", ""),
		70:  newAdSystem(70, "OMS", ""),
		71:  newAdSystem(71, "Ströer", ""),
		73:  newAdSystem(73, "C1X", ""),
		74:  newAdSystem(74, "Synacor", ""),
		76:  newAdSystem(76, "Videology", ""),
		77:  newAdSystem(77, "Telaria (fka Tremor Video)", "tremorhub.com"),
		78:  newAdSystem(78, "Genesis Media", "altitude-arena.com"),
		80:  newAdSystem(80, "Imonomy", ""),
		81:  newAdSystem(81, "Komoona", ""),
		82:  newAdSystem(82, "SpringServe", ""),
		83:  newAdSystem(83, "TripleLift", ""),
		84:  newAdSystem(84, "AppNexus", "appnexus.com"),
		85:  newAdSystem(85, "NTV", ""),
		86:  newAdSystem(86, "COMET", ""),
		87:  newAdSystem(87, "Undertone", ""),
		88:  newAdSystem(88, "One by AOL: Video", "advertising.com"),
		89:  newAdSystem(89, "Algovid", ""),
		90:  newAdSystem(90, "Lockerdome", ""),
		91:  newAdSystem(91, "Widespace", ""),
		92:  newAdSystem(92, "Sortable", ""),
		93:  newAdSystem(93, "Mobfox", ""),
		94:  newAdSystem(94, "Teads", "teads.tv"),
		95:  newAdSystem(95, "PulsePoint", "contextweb.com"),
		96:  newAdSystem(96, "District M", ""),
		97:  newAdSystem(97, "Sharethrough", ""),
		98:  newAdSystem(98, "Adfrontiers", ""),
		99:  newAdSystem(99, "Ad3media", ""),
		100: newAdSystem(100, "ADMIZED", ""),
		101: newAdSystem(101, "Twiago", ""),
		102: newAdSystem(102, "Xapads", ""),
		104: newAdSystem(104, "Adstir", ""),
		105: newAdSystem(105, "Yieldlab", ""),
		107: newAdSystem(107, "Ad6Media", ""),
		108: newAdSystem(108, "Adbistro", ""),
		109: newAdSystem(109, "AdColony", ""),
		110: newAdSystem(110, "Fluct", ""),
		111: newAdSystem(111, "Adman Media", ""),
		112: newAdSystem(112, "AdMedia", ""),
		113: newAdSystem(113, "AdMixer", ""),
		114: newAdSystem(114, "NOT IN USE", ""),
		115: newAdSystem(115, "Ads4Pics", ""),
		117: newAdSystem(117, "Adunity", ""),
		118: newAdSystem(118, "AMM Media Marketing", ""),
		119: newAdSystem(119, "Advertise.com", ""),
		120: newAdSystem(120, "Aerserv", ""),
		121: newAdSystem(121, "AndBeyond.Media", ""),
		122: newAdSystem(122, "appTV", ""),
		123: newAdSystem(123, "ucfunnel", ""),
		124: newAdSystem(124, "WideOrbit", ""),
		125: newAdSystem(125, "Aximus", ""),
		126: newAdSystem(126, "BaronsMedia", ""),
		128: newAdSystem(128, "Streamlyn", ""),
		129: newAdSystem(129, "Bidtheater", ""),
		131: newAdSystem(131, "Buy Sell Ads", ""),
		132: newAdSystem(132, "Carambola", ""),
		133: newAdSystem(133, "Cedato", ""),
		134: newAdSystem(134, "Clickio", ""),
		135: newAdSystem(135, "Collective", ""),
		136: newAdSystem(136, "Adimia", ""),
		137: newAdSystem(137, "Converge-Digital", ""),
		138: newAdSystem(138, "Crimtan", ""),
		139: newAdSystem(139, "Defy", ""),
		141: newAdSystem(141, "DistroScale", ""),
		142: newAdSystem(142, "DynAdmic", ""),
		144: newAdSystem(144, "EADV", ""),
		145: newAdSystem(145, "Easy Platform", ""),
		146: newAdSystem(146, "eBoundServices", ""),
		147: newAdSystem(147, "Electric Sheep", ""),
		148: newAdSystem(148, "FirstImpression.io", ""),
		149: newAdSystem(149, "Exclude", ""),
		150: newAdSystem(150, "Get Intent", ""),
		151: newAdSystem(151, "Glu Company", ""),
		152: newAdSystem(152, "GMO SSP", ""),
		153: newAdSystem(153, "Browsi", ""),
		154: newAdSystem(154, "Gourmet Ads", ""),
		155: newAdSystem(155, "Hiro Media", ""),
		156: newAdSystem(156, "iBillboard", ""),
		157: newAdSystem(157, "Increase Rev", ""),
		158: newAdSystem(158, "Infolinks", ""),
		159: newAdSystem(159, "Insticator", ""),
		160: newAdSystem(160, "JustPremium", ""),
		161: newAdSystem(161, "JWPlayer", ""),
		162: newAdSystem(162, "KeenKale", ""),
		163: newAdSystem(163, "Lifestreet", ""),
		164: newAdSystem(164, "Linicom", ""),
		165: newAdSystem(165, "MadAdsMedia", ""),
		166: newAdSystem(166, "Vuble", "mediabong.com"),
		167: newAdSystem(167, "Deguate", ""),
		169: newAdSystem(169, "Mgid", ""),
		170: newAdSystem(170, "Monarch Ads", ""),
		171: newAdSystem(171, "Netseer", ""),
		173: newAdSystem(173, "Ooyala", ""),
		174: newAdSystem(174, "Optimatic", ""),
		175: newAdSystem(175, "Padsquad", ""),
		176: newAdSystem(176, "Paypal", ""),
		177: newAdSystem(177, "Playtouch", ""),
		178: newAdSystem(178, "Paywire", ""),
		179: newAdSystem(179, "PowerLinks", ""),
		180: newAdSystem(180, "NexTag", ""),
		181: newAdSystem(181, "Purch", ""),
		182: newAdSystem(182, "Q1 Media", ""),
		183: newAdSystem(183, "Quantcast", ""),
		184: newAdSystem(184, "Quantum Native", ""),
		185: newAdSystem(185, "ReklamStore", ""),
		186: newAdSystem(186, "RekMob", ""),
		188: newAdSystem(188, "Smartclip", ""),
		189: newAdSystem(189, "Smarty Ads", ""),
		190: newAdSystem(190, "Somo Audience", "somoaudience.com"),
		191: newAdSystem(191, "Spot.im", ""),
		192: newAdSystem(192, "Sprout", ""),
		193: newAdSystem(193, "SSPHwy", ""),
		194: newAdSystem(194, "StartApp", ""),
		195: newAdSystem(195, "SNT Media", ""),
		196: newAdSystem(196, "TabletMedia", ""),
		197: newAdSystem(197, "Tappx", ""),
		198: newAdSystem(198, "The Moneytizer", ""),
		199: newAdSystem(199, "The Trade Desk", ""),
		200: newAdSystem(200, "Thrive", ""),
		201: newAdSystem(201, "Tisoomi", ""),
		202: newAdSystem(202, "Tribal Fusion", ""),
		203: newAdSystem(203, "Trion Interactive", ""),
		204: newAdSystem(204, "TrueX", ""),
		205: newAdSystem(205, "Turf Digital", ""),
		206: newAdSystem(206, "UBM", ""),
		207: newAdSystem(207, "Underdog Media", ""),
		208: newAdSystem(208, "Alliance Data", ""),
		209: newAdSystem(209, "Verta Media", ""),
		210: newAdSystem(210, "Vertoz", ""),
		211: newAdSystem(211, "Video Intelligence", ""),
		212: newAdSystem(212, "Fidelity Media", ""),
		213: newAdSystem(213, "Yandex", ""),
		214: newAdSystem(214, "Yellow Hammer", ""),
		215: newAdSystem(215, "RockYou", "rockyou.net"),
		216: newAdSystem(216, "Innity", "innity.com"),
		217: newAdSystem(217, "Native Ads", "nativeads.com"),
		218: newAdSystem(218, "RichAudience", "richaudience.com"),
		219: newAdSystem(219, "AdStanding", "adstanding.com"),
		220: newAdSystem(220, "Mass2", "www.mass2.com"),
		221: newAdSystem(221, "RTK.io", ""),
		222: newAdSystem(222, "Atomx", "atomx.com"),
		223: newAdSystem(223, "Addroplet.com ", "Addroplet.com "),
		224: newAdSystem(224, "Liondigitalserving.com", "Liondigitalserving.com"),
		225: newAdSystem(225, "sulvo.com", "sulvo.com"),
		226: newAdSystem(226, "surgeprice.com", "surgeprice.com"),
		227: newAdSystem(227, "mediabong.com", "mediabong.com"),
		228: newAdSystem(228, "Seracast", "babaroll.com"),
		229: newAdSystem(229, "Juice Nectar", "juicenectar.com"),
		230: newAdSystem(230, "AdPone", "adpone.com"),
		231: newAdSystem(231, "OneTag", "onetag.com"),
		232: newAdSystem(232, "Between Exchange", "betweendigital.com"),
		233: newAdSystem(233, "Experian", "experian.com"),
		234: newAdSystem(234, "GammaSSP", "gammassp.com"),
		235: newAdSystem(235, "Cynogage", "cynogage.com"),
		236: newAdSystem(236, "DeepIntent", "deepintent.com"),
		237: newAdSystem(237, "Adversal", "adversal.com"),
		238: newAdSystem(238, "vmg.host", "vmg.host"),
		239: newAdSystem(239, "Vdopia", ""),
		240: newAdSystem(240, "Yengo", ""),
		241: newAdSystem(241, "Backbeatmedia", ""),
		242: newAdSystem(242, "Videmob by Cydersoft", ""),
		243: newAdSystem(243, "Ligatus", ""),
		244: newAdSystem(244, "Vidstart", ""),
	}

	adSystemDomains = map[string]*adSystemDomain{
		"rubicon.com":                                   newAdSystemDomain("rubicon.com", 1),
		"fastlane.rubiconproject.com":                   newAdSystemDomain("fastlane.rubiconproject.com", 1),
		"ads.rubiconproject.com":                        newAdSystemDomain("ads.rubiconproject.com", 1),
		"rubiconproject.com":                            newAdSystemDomain("rubiconproject.com", 1),
		"rubiconproject.com<http://rubiconproject.com>": newAdSystemDomain("rubiconproject.com<http://rubiconproject.com>", 1),
		"33across.com":                                  newAdSystemDomain("33across.com", 2),
		"pubmatic.com":                                  newAdSystemDomain("pubmatic.com", 3),
		"apps.pubmatic.com":                             newAdSystemDomain("apps.pubmatic.com", 3),
		"pubmatic":                                      newAdSystemDomain("pubmatic", 3),
		"openx.com":                                     newAdSystemDomain("openx.com", 4),
		"openx":                                         newAdSystemDomain("openx", 4),
		"openxebda":                                     newAdSystemDomain("openxebda", 4),
		"openxprebid":                                   newAdSystemDomain("openxprebid", 4),
		"openx.com<http://openx.com>":                   newAdSystemDomain("openx.com<http://openx.com>", 4),
		"openx.net":                                     newAdSystemDomain("openx.net", 4),
		"facebook.com":                                  newAdSystemDomain("facebook.com", 5),
		"facebook":                                      newAdSystemDomain("facebook", 5),
		"facebook:facebook.com":                         newAdSystemDomain("facebook:facebook.com", 5),
		"gumgum.com":                                    newAdSystemDomain("gumgum.com", 6),
		"kargo.com":                                     newAdSystemDomain("kargo.com", 7),
		"google.com":                                    newAdSystemDomain("google.com", 8),
		"googletagservices.com":                         newAdSystemDomain("googletagservices.com", 8),
		"?google.com":                                   newAdSystemDomain("?google.com", 8),
		"adsense":                                       newAdSystemDomain("adsense", 8),
		"google.com/adsense":                            newAdSystemDomain("google.com/adsense", 8),
		"google.com<http://google.com>":                 newAdSystemDomain("google.com<http://google.com>", 8),
		"www.google.com/dfp":                            newAdSystemDomain("www.google.com/dfp", 8),
		"brealtime.com":                                 newAdSystemDomain("brealtime.com", 9),
		"Brealtime":                                     newAdSystemDomain("Brealtime", 9),
		"brealtimegoogle":                               newAdSystemDomain("brealtimegoogle", 9),
		"emxdgt.com105":                                 newAdSystemDomain("emxdgt.com105", 9),
		"amazon-adsystem.com":                           newAdSystemDomain("amazon-adsystem.com", 10),
		"c.amazon-adsystem.com":                         newAdSystemDomain("c.amazon-adsystem.com", 10),
		"advertising.amazon.com":                        newAdSystemDomain("advertising.amazon.com", 10),
		"amazon.com":                                    newAdSystemDomain("amazon.com", 10),
		"a9.com":                                        newAdSystemDomain("a9.com", 10),
		"aps.amazon.com":                                newAdSystemDomain("aps.amazon.com", 10),
		"adtech.com":                                    newAdSystemDomain("adtech.com", 11),
		"adtech.net":                                    newAdSystemDomain("adtech.net", 11),
		"aolcloud.net":                                  newAdSystemDomain("aolcloud.net", 11),
		"liveintent.com":                                newAdSystemDomain("liveintent.com", 12),
		"yieldmo.com":                                   newAdSystemDomain("yieldmo.com", 13),
		"mopub.com":                                     newAdSystemDomain("mopub.com", 14),
		"aol.com":                                       newAdSystemDomain("aol.com", 15),
		"smartstream.tv":                                newAdSystemDomain("smartstream.tv", 16),
		"smaato.com":                                    newAdSystemDomain("smaato.com", 17),
		"spx.smaato.com":                                newAdSystemDomain("spx.smaato.com", 17),
		"taboola.com":                                   newAdSystemDomain("taboola.com", 18),
		"trustx.org":                                    newAdSystemDomain("trustx.org", 19),
		"sofia.trustx.org":                              newAdSystemDomain("sofia.trustx.org", 19),
		"lkqd.net":                                      newAdSystemDomain("lkqd.net", 20),
		"lkqd.com":                                      newAdSystemDomain("lkqd.com", 20),
		"ad.lkqd.net":                                   newAdSystemDomain("ad.lkqd.net", 20),
		"criteo.com":                                    newAdSystemDomain("criteo.com", 21),
		"critero.com":                                   newAdSystemDomain("critero.com", 21),
		"criteo.net":                                    newAdSystemDomain("criteo.net", 21),
		"phillymag.com==criteo.com":                     newAdSystemDomain("phillymag.com==criteo.com", 21),
		"exponential.com":                               newAdSystemDomain("exponential.com", 22),
		"exponential.comi":                              newAdSystemDomain("exponential.comi", 22),
		"xponential.com":                                newAdSystemDomain("xponential.com", 22),
		"lijit.com":                                     newAdSystemDomain("lijit.com", 23),
		"meridian.sovrn.com":                            newAdSystemDomain("meridian.sovrn.com", 23),
		"sovrn.com":                                     newAdSystemDomain("sovrn.com", 23),
		"lijit":                                         newAdSystemDomain("lijit", 23),
		"rhythmone.com":                                 newAdSystemDomain("rhythmone.com", 24),
		"1rx.io":                                        newAdSystemDomain("1rx.io", 24),
		"yldbt.com":                                     newAdSystemDomain("yldbt.com", 25),
		"technorati.com":                                newAdSystemDomain("technorati.com", 26),
		"bidfluence.com":                                newAdSystemDomain("bidfluence.com", 27),
		"beachfront.com":                                newAdSystemDomain("beachfront.com", 27),
		"switch.com":                                    newAdSystemDomain("switch.com", 28),
		"switchconcept":                                 newAdSystemDomain("switchconcept", 28),
		"switchconcepts.com":                            newAdSystemDomain("switchconcepts.com", 28),
		"brightroll.com":                                newAdSystemDomain("brightroll.com", 29),
		"conversantmedia.com":                           newAdSystemDomain("conversantmedia.com", 30),
		"go.sonobi.com":                                 newAdSystemDomain("go.sonobi.com", 31),
		"sonobi.com":                                    newAdSystemDomain("sonobi.com", 31),
		"*.go.sonobi.com":                               newAdSystemDomain("*.go.sonobi.com", 31),
		"spoutable.com":                                 newAdSystemDomain("spoutable.com", 32),
		"freewheel.tv":                                  newAdSystemDomain("freewheel.tv", 33),
		"cdn.stickyadstv.com":                           newAdSystemDomain("cdn.stickyadstv.com", 33),
		"stickyad:freewheel.tv":                         newAdSystemDomain("stickyad:freewheel.tv", 33),
		"connatix.com":                                  newAdSystemDomain("connatix.com", 34),
		"t.brand-server.com":                            newAdSystemDomain("t.brand-server.com", 35),
		"positivemobile.com":                            newAdSystemDomain("positivemobile.com", 36),
		"memeglobal.com":                                newAdSystemDomain("memeglobal.com", 37),
		"kixer.com":                                     newAdSystemDomain("kixer.com", 38),
		"sekindo.com":                                   newAdSystemDomain("sekindo.com", 39),
		"sekindo":                                       newAdSystemDomain("sekindo", 39),
		"360yield.com":                                  newAdSystemDomain("360yield.com", 40),
		"improvedigital.com":                            newAdSystemDomain("improvedigital.com", 40),
		"adform.com":                                    newAdSystemDomain("adform.com", 41),
		"adform.net":                                    newAdSystemDomain("adform.net", 41),
		"adx.adform.net":                                newAdSystemDomain("adx.adform.net", 41),
		"inner-active.com":                              newAdSystemDomain("inner-active.com", 43),
		"spotxchange.com":                               newAdSystemDomain("spotxchange.com", 44),
		"spotx.tv":                                      newAdSystemDomain("spotx.tv", 44),
		"streamrail.net":                                newAdSystemDomain("streamrail.net", 45),
		"sdk.streamrail.com":                            newAdSystemDomain("sdk.streamrail.com", 45),
		"mathtag.com":                                   newAdSystemDomain("mathtag.com", 46),
		"mediamath.com":                                 newAdSystemDomain("mediamath.com", 46),
		"adyoulike.com":                                 newAdSystemDomain("adyoulike.com", 47),
		"indexexchnage.com":                             newAdSystemDomain("indexexchnage.com", 48),
		"indexexchange.com":                             newAdSystemDomain("indexexchange.com", 48),
		"www.indexexchange.com":                         newAdSystemDomain("www.indexexchange.com", 48),
		"indexechange.com":                              newAdSystemDomain("indexechange.com", 48),
		"indexexchange(ebda)":                           newAdSystemDomain("indexexchange(ebda)", 48),
		"indexexchange(pubmatic)":                       newAdSystemDomain("indexexchange(pubmatic)", 48),
		"indexexchange(videossp)":                       newAdSystemDomain("indexexchange(videossp)", 48),
		"index.com":                                     newAdSystemDomain("index.com", 48),
		"kiosked.com":                                   newAdSystemDomain("kiosked.com", 50),
		"ads.kiosked.com":                               newAdSystemDomain("ads.kiosked.com", 50),
		"video.unrulymedia.com":                         newAdSystemDomain("video.unrulymedia.com", 51),
		"brightcom.com":                                 newAdSystemDomain("brightcom.com", 52),
		"rs-stripe.com":                                 newAdSystemDomain("rs-stripe.com", 53),
		"fyber.com":                                     newAdSystemDomain("fyber.com", 54),
		"tidaltv.com":                                   newAdSystemDomain("tidaltv.com", 55),
		"nativo.com":                                    newAdSystemDomain("nativo.com", 56),
		"jadserve.postrelease.com":                      newAdSystemDomain("jadserve.postrelease.com", 56),
		"media.net":                                     newAdSystemDomain("media.net", 57),
		"www.yumenetworks.com":                          newAdSystemDomain("www.yumenetworks.com", 58),
		"yume.com":                                      newAdSystemDomain("yume.com", 58),
		"yumenetworks.com":                              newAdSystemDomain("yumenetworks.com", 58),
		"revcontent.com":                                newAdSystemDomain("revcontent.com", 59),
		"revontent.com":                                 newAdSystemDomain("revontent.com", 59),
		"outbrain.com":                                  newAdSystemDomain("outbrain.com", 60),
		"zedo.com":                                      newAdSystemDomain("zedo.com", 61),
		"freeskreen.com":                                newAdSystemDomain("freeskreen.com", 62),
		"bidtellect.com":                                newAdSystemDomain("bidtellect.com", 63),
		"smartadserver.com":                             newAdSystemDomain("smartadserver.com", 64),
		"loopme.com":                                    newAdSystemDomain("loopme.com", 65),
		"vidazoo.com":                                   newAdSystemDomain("vidazoo.com", 66),
		"vidazoo":                                       newAdSystemDomain("vidazoo", 66),
		"videoflare.com":                                newAdSystemDomain("videoflare.com", 67),
		"yahoo.com":                                     newAdSystemDomain("yahoo.com", 68),
		"pixfuture.com":                                 newAdSystemDomain("pixfuture.com", 69),
		"oms.eu":                                        newAdSystemDomain("oms.eu", 70),
		"stroeer.com":                                   newAdSystemDomain("stroeer.com", 71),
		"emxdgt.com":                                    newAdSystemDomain("emxdgt.com", 9),
		"c1exchange.com":                                newAdSystemDomain("c1exchange.com", 73),
		"synacor.com":                                   newAdSystemDomain("synacor.com", 74),
		"sfx.freewheel.tv":                              newAdSystemDomain("sfx.freewheel.tv", 33),
		"videologygroup.com":                            newAdSystemDomain("videologygroup.com", 76),
		"tremorhub.com":                                 newAdSystemDomain("tremorhub.com", 77),
		"altitudedigital.com":                           newAdSystemDomain("altitudedigital.com", 78),
		"platform.videologygroup.com":                   newAdSystemDomain("platform.videologygroup.com", 76),
		"imonomy.com":                                   newAdSystemDomain("imonomy.com", 80),
		"komoona ltd":                                   newAdSystemDomain("komoona ltd", 81),
		"komoonaltd":                                    newAdSystemDomain("komoonaltd", 81),
		"springserve.com":                               newAdSystemDomain("springserve.com", 82),
		"spingserve.com":                                newAdSystemDomain("spingserve.com", 82),
		"triplelift.com":                                newAdSystemDomain("triplelift.com", 83),
		"www.triplelift.com":                            newAdSystemDomain("www.triplelift.com", 83),
		"ib.adnxs.com":                                  newAdSystemDomain("ib.adnxs.com", 84),
		"appnexus.com":                                  newAdSystemDomain("appnexus.com", 84),
		"appnexus":                                      newAdSystemDomain("appnexus", 84),
		"apnexus.com":                                   newAdSystemDomain("apnexus.com", 84),
		"appnexus.txt":                                  newAdSystemDomain("appnexus.txt", 84),
		"adnxs.com":                                     newAdSystemDomain("adnxs.com", 84),
		"appnexus.com<http://appnexus.com>":             newAdSystemDomain("appnexus.com<http://appnexus.com>", 84),
		"s.ntv.io/serve":                                newAdSystemDomain("s.ntv.io/serve", 85),
		"coxmt.com":                                     newAdSystemDomain("coxmt.com", 86),
		"undertone.com":                                 newAdSystemDomain("undertone.com", 87),
		"advertising.com":                               newAdSystemDomain("advertising.com", 88),
		"c.algovid.com":                                 newAdSystemDomain("c.algovid.com", 89),
		"lockerdome.com":                                newAdSystemDomain("lockerdome.com", 90),
		"widespace.com":                                 newAdSystemDomain("widespace.com", 91),
		"deployads.com":                                 newAdSystemDomain("deployads.com", 92),
		"www.mobfox.com":                                newAdSystemDomain("www.mobfox.com", 93),
		"mobfox.com":                                    newAdSystemDomain("mobfox.com", 93),
		"teads.tv":                                      newAdSystemDomain("teads.tv", 94),
		"teads.com":                                     newAdSystemDomain("teads.com", 94),
		"publishers.teads.tv":                           newAdSystemDomain("publishers.teads.tv", 94),
		"contextweb.com":                                newAdSystemDomain("contextweb.com", 95),
		"pulsepoint.com":                                newAdSystemDomain("pulsepoint.com", 95),
		"pulsepoint":                                    newAdSystemDomain("pulsepoint", 95),
		"pulsepoint:contextweb.com":                     newAdSystemDomain("pulsepoint:contextweb.com", 95),
		"districtm.com":                                 newAdSystemDomain("districtm.com", 96),
		"districtm.ca":                                  newAdSystemDomain("districtm.ca", 96),
		"districtm.io":                                  newAdSystemDomain("districtm.io", 96),
		"sharethrough.com":                              newAdSystemDomain("sharethrough.com", 97),
		"media.adfrontiers.com":                         newAdSystemDomain("media.adfrontiers.com", 98),
		"adfrontiers.com":                               newAdSystemDomain("adfrontiers.com", 98),
		"media.adfrontiers":                             newAdSystemDomain("media.adfrontiers", 98),
		"ad3media.com":                                  newAdSystemDomain("ad3media.com", 99),
		"ads.admized.com":                               newAdSystemDomain("ads.admized.com", 100),
		"admized.com":                                   newAdSystemDomain("admized.com", 100),
		"a.twiago.com":                                  newAdSystemDomain("a.twiago.com", 101),
		"twiago.com":                                    newAdSystemDomain("twiago.com", 101),
		"xapads.com":                                    newAdSystemDomain("xapads.com", 102),
		"ad-stir.com":                                   newAdSystemDomain("ad-stir.com", 104),
		"ad.yieldlab.net":                               newAdSystemDomain("ad.yieldlab.net", 105),
		"yieldlab.de":                                   newAdSystemDomain("yieldlab.de", 105),
		"yieldlab.net":                                  newAdSystemDomain("yieldlab.net", 105),
		"ad3.io":                                        newAdSystemDomain("ad3.io", 99),
		"ad6media.es":                                   newAdSystemDomain("ad6media.es", 107),
		"ad6media.fr":                                   newAdSystemDomain("ad6media.fr", 107),
		"www.ad6media.fr":                               newAdSystemDomain("www.ad6media.fr", 107),
		"adbistro.com":                                  newAdSystemDomain("adbistro.com", 108),
		"adcolony.com":                                  newAdSystemDomain("adcolony.com", 109),
		"adingo.jp":                                     newAdSystemDomain("adingo.jp", 110),
		"adingo.jp<http://adingo.jp>":                   newAdSystemDomain("adingo.jp<http://adingo.jp>", 110),
		"admanmedia.com":                                newAdSystemDomain("admanmedia.com", 111),
		"admedia.com":                                   newAdSystemDomain("admedia.com", 112),
		"admixer.com":                                   newAdSystemDomain("admixer.com", 113),
		"admixer.net":                                   newAdSystemDomain("admixer.net", 113),
		"ads.stickyadstv.com":                           newAdSystemDomain("ads.stickyadstv.com", 33),
		"ads4pics.com":                                  newAdSystemDomain("ads4pics.com", 115),
		"adtech.com<http://adtech.com>":                 newAdSystemDomain("adtech.com<http://adtech.com>", 11),
		"aolcloud.com":                                  newAdSystemDomain("aolcloud.com", 11),
		"aolcloud.net<http://aolcloud.net>":             newAdSystemDomain("aolcloud.net<http://aolcloud.net>", 11),
		"adunity.com":                                   newAdSystemDomain("adunity.com", 117),
		"advbo.ammadv.it":                               newAdSystemDomain("advbo.ammadv.it", 118),
		"Advertise.com":                                 newAdSystemDomain("Advertise.com", 119),
		"advertising.com<http://advertising.com>": newAdSystemDomain("advertising.com<http://advertising.com>", 88),
		"aerserv.com":                     newAdSystemDomain("aerserv.com", 120),
		"andbeyond.media":                 newAdSystemDomain("andbeyond.media", 121),
		"app.tv":                          newAdSystemDomain("app.tv", 122),
		"apptv.com":                       newAdSystemDomain("apptv.com", 122),
		"aralego.com":                     newAdSystemDomain("aralego.com", 123),
		"atemda.com":                      newAdSystemDomain("atemda.com", 124),
		"aximusag":                        newAdSystemDomain("aximusag", 125),
		"aximus.ch":                       newAdSystemDomain("aximus.ch", 125),
		"baronsmedia.com":                 newAdSystemDomain("baronsmedia.com", 126),
		"bidsxchange.com":                 newAdSystemDomain("bidsxchange.com", 128),
		"bidtheatre.com":                  newAdSystemDomain("bidtheatre.com", 129),
		"buysellads.com":                  newAdSystemDomain("buysellads.com", 131),
		"carambo.la":                      newAdSystemDomain("carambo.la", 132),
		"carambola.com":                   newAdSystemDomain("carambola.com", 132),
		"cedato.com":                      newAdSystemDomain("cedato.com", 133),
		"clickio.com":                     newAdSystemDomain("clickio.com", 134),
		"collectiveuk.com":                newAdSystemDomain("collectiveuk.com", 135),
		"connectignite.com":               newAdSystemDomain("connectignite.com", 136),
		"converge-digital.com":            newAdSystemDomain("converge-digital.com", 137),
		"crimtan.com":                     newAdSystemDomain("crimtan.com", 138),
		"defymedia.com":                   newAdSystemDomain("defymedia.com", 139),
		"distrcitm.io":                    newAdSystemDomain("distrcitm.io", 96),
		"districtmadexchange":             newAdSystemDomain("districtmadexchange", 96),
		"districtm":                       newAdSystemDomain("districtm", 96),
		"districtm.net":                   newAdSystemDomain("districtm.net", 96),
		"districtmio.com":                 newAdSystemDomain("districtmio.com", 96),
		"distroscale.com":                 newAdSystemDomain("distroscale.com", 141),
		"dynadmic":                        newAdSystemDomain("dynadmic", 142),
		"e-planning.net":                  newAdSystemDomain("e-planning.net", 49),
		"eadv.it":                         newAdSystemDomain("eadv.it", 144),
		"easyplatform.com":                newAdSystemDomain("easyplatform.com", 145),
		"eboundservices.com":              newAdSystemDomain("eboundservices.com", 146),
		"electric-sheep.tv":               newAdSystemDomain("electric-sheep.tv", 147),
		"firstimpression.io":              newAdSystemDomain("firstimpression.io", 148),
		"geekexchange.com":                newAdSystemDomain("geekexchange.com", 149),
		"getintent.com":                   newAdSystemDomain("getintent.com", 150),
		"glucompany.com":                  newAdSystemDomain("glucompany.com", 151),
		"gmossp.jp":                       newAdSystemDomain("gmossp.jp", 152),
		"gobrowsi.com":                    newAdSystemDomain("gobrowsi.com", 153),
		"gourmetads.com":                  newAdSystemDomain("gourmetads.com", 154),
		"hiro-media.com":                  newAdSystemDomain("hiro-media.com", 155),
		"ibillboard.com":                  newAdSystemDomain("ibillboard.com", 156),
		"increaserev.com":                 newAdSystemDomain("increaserev.com", 157),
		"infolinks.com":                   newAdSystemDomain("infolinks.com", 158),
		"insticator.com":                  newAdSystemDomain("insticator.com", 159),
		"justpremium.com":                 newAdSystemDomain("justpremium.com", 160),
		"jwdemandadexchange":              newAdSystemDomain("jwdemandadexchange", 161),
		"keenkale.com":                    newAdSystemDomain("keenkale.com", 162),
		"lifestreet.com":                  newAdSystemDomain("lifestreet.com", 163),
		"linicom":                         newAdSystemDomain("linicom", 164),
		"madadsmedia.com":                 newAdSystemDomain("madadsmedia.com", 165),
		"mediabong.net":                   newAdSystemDomain("mediabong.net", 166),
		"mediadeguate.com":                newAdSystemDomain("mediadeguate.com", 167),
		"memevideoad.com":                 newAdSystemDomain("memevideoad.com", 37),
		"stinger.memeglobal.com":          newAdSystemDomain("stinger.memeglobal.com", 37),
		"mgid.com":                        newAdSystemDomain("mgid.com", 169),
		"monarchads.com":                  newAdSystemDomain("monarchads.com", 170),
		"netseer.com":                     newAdSystemDomain("netseer.com", 171),
		"oogle.com":                       newAdSystemDomain("oogle.com", 8),
		"ooyala.com":                      newAdSystemDomain("ooyala.com", 173),
		"optimatic.com":                   newAdSystemDomain("optimatic.com", 174),
		"padsquad.com":                    newAdSystemDomain("padsquad.com", 175),
		"paypal.com":                      newAdSystemDomain("paypal.com", 176),
		"playtouch":                       newAdSystemDomain("playtouch", 177),
		"playtouch2":                      newAdSystemDomain("playtouch2", 177),
		"playwire.com":                    newAdSystemDomain("playwire.com", 178),
		"powerlinks.com":                  newAdSystemDomain("powerlinks.com", 179),
		"pubgears.com":                    newAdSystemDomain("pubgears.com", 180),
		"purch.com":                       newAdSystemDomain("purch.com", 181),
		"servebom.com":                    newAdSystemDomain("servebom.com", 181),
		"q1media.com":                     newAdSystemDomain("q1media.com", 182),
		"q1connect.com":                   newAdSystemDomain("q1connect.com", 182),
		"quantcast.com":                   newAdSystemDomain("quantcast.com", 183),
		"quantum-advertising.com":         newAdSystemDomain("quantum-advertising.com", 184),
		"reklamstore.com":                 newAdSystemDomain("reklamstore.com", 185),
		"rekmob.com":                      newAdSystemDomain("rekmob.com", 186),
		"smartadserver:smartadserver.com": newAdSystemDomain("smartadserver:smartadserver.com", 64),
		"smartadsever.com":                newAdSystemDomain("smartadsever.com", 64),
		"smartclip.net":                   newAdSystemDomain("smartclip.net", 188),
		"smartyads.com":                   newAdSystemDomain("smartyads.com", 189),
		"somoaudience.com":                newAdSystemDomain("somoaudience.com", 190),
		"SpotIM":                          newAdSystemDomain("SpotIM", 191),
		"sprout-ad.com":                   newAdSystemDomain("sprout-ad.com", 192),
		"ssphwy.com":                      newAdSystemDomain("ssphwy.com", 193),
		"startapp.com":                    newAdSystemDomain("startapp.com", 194),
		"synapsys.us":                     newAdSystemDomain("synapsys.us", 195),
		"tabletmedia.co.uk":               newAdSystemDomain("tabletmedia.co.uk", 196),
		"tappx.com":                       newAdSystemDomain("tappx.com", 197),
		"themoneytizer.com":               newAdSystemDomain("themoneytizer.com", 198),
		"thetradedesk.com":                newAdSystemDomain("thetradedesk.com", 199),
		"thrive.plus":                     newAdSystemDomain("thrive.plus", 200),
		"tisoomi-services.com":            newAdSystemDomain("tisoomi-services.com", 201),
		"tribalfusion.com":                newAdSystemDomain("tribalfusion.com", 202),
		"trion.com":                       newAdSystemDomain("trion.com", 203),
		"trioninteractive.com":            newAdSystemDomain("trioninteractive.com", 203),
		"truex.com":                       newAdSystemDomain("truex.com", 204),
		"turf.digital":                    newAdSystemDomain("turf.digital", 205),
		"ubm.com":                         newAdSystemDomain("ubm.com", 206),
		"udmserve.net":                    newAdSystemDomain("udmserve.net", 207),
		"valueclickmedia.com":             newAdSystemDomain("valueclickmedia.com", 208),
		"vertamedia.com":                  newAdSystemDomain("vertamedia.com", 209),
		"vertoz.com":                      newAdSystemDomain("vertoz.com", 210),
		"vi.ai":                           newAdSystemDomain("vi.ai", 211),
		"www.vi.ai":                       newAdSystemDomain("www.vi.ai", 211),
		"x.fidelity-media.com":            newAdSystemDomain("x.fidelity-media.com", 212),
		"yandex.ru":                       newAdSystemDomain("yandex.ru", 213),
		"yellowhammer.com":                newAdSystemDomain("yellowhammer.com", 214),
		"rockyou.com":                     newAdSystemDomain("rockyou.com", 215),
		"rockyou.net":                     newAdSystemDomain("rockyou.net", 215),
		"innity.com":                      newAdSystemDomain("innity.com", 216),
		"innity.net":                      newAdSystemDomain("innity.net", 216),
		"advenueplatform.com":             newAdSystemDomain("advenueplatform.com", 216),
		"nativeads.com":                   newAdSystemDomain("nativeads.com", 217),
		"natiiveads.com":                  newAdSystemDomain("natiiveads.com", 217),
		"richaudience.com":                newAdSystemDomain("richaudience.com", 218),
		"adstanding.com":                  newAdSystemDomain("adstanding.com", 219),
		"www.mass2.com":                   newAdSystemDomain("www.mass2.com", 220),
		"RTK.io":                          newAdSystemDomain("RTK.io", 221),
		"atomx.com":                       newAdSystemDomain("atomx.com", 222),
		"ato.mx":                          newAdSystemDomain("ato.mx", 222),
		"rtb.ato.mx":                      newAdSystemDomain("rtb.ato.mx", 222),
		"p.ato.mx":                        newAdSystemDomain("p.ato.mx", 222),
		"addroplet.com":                   newAdSystemDomain("addroplet.com", 223),
		"Liondigitalserving.com":          newAdSystemDomain("Liondigitalserving.com", 224),
		"sulvo.com":                       newAdSystemDomain("sulvo.com", 225),
		"surgeprice.com":                  newAdSystemDomain("surgeprice.com", 226),
		"mediabong.com":                   newAdSystemDomain("mediabong.com", 227),
		"babaroll.com":                    newAdSystemDomain("babaroll.com", 228),
		"juicenectar.com":                 newAdSystemDomain("juicenectar.com", 229),
		"adpone.com":                      newAdSystemDomain("adpone.com", 230),
		"onetag.com":                      newAdSystemDomain("onetag.com", 231),
		"onetag-sys.com":                  newAdSystemDomain("onetag-sys.com", 231),
		"betweendigital.com":              newAdSystemDomain("betweendigital.com", 232),
		"ads.betweendigital.com":          newAdSystemDomain("ads.betweendigital.com", 232),
		"experian.com":                    newAdSystemDomain("experian.com", 233),
		"altitude-arena.com":              newAdSystemDomain("altitude-arena.com", 78),
		"gammassp.com":                    newAdSystemDomain("gammassp.com", 234),
		"ambientdigitalgroup.com":         newAdSystemDomain("ambientdigitalgroup.com", 234),
		"cynogage.com":                    newAdSystemDomain("cynogage.com", 235),
		"deepintent.com":                  newAdSystemDomain("deepintent.com", 236),
		"adversal.com":                    newAdSystemDomain("adversal.com", 237),
		"vmg.host":                        newAdSystemDomain("vmg.host", 238),
		"Chocolateplatform.com":           newAdSystemDomain("Chocolateplatform.com", 239),
		"directadvert.ru":                 newAdSystemDomain("directadvert.ru", 240),
		"backbeatmedia.com":               newAdSystemDomain("backbeatmedia.com", 241),
		"videmob.com":                     newAdSystemDomain("videmob.com", 242),
		"ligadx.com":                      newAdSystemDomain("ligadx.com", 243),
		"vidstart.com":                    newAdSystemDomain("vidstart.com", 244),
		"mobileadtrading.com":             newAdSystemDomain("mobileadtrading.com", 245),
	}
}

// adSystem single know ad system (SSPs/exchanges).
type adSystem struct {
	ID              int    // ID of the ad system: there is no order or meaning implied by the ID, it is merely an auto incrementing number
	Name            string // Name holds the name of the AdSystem
	CanonicalDomain string // CanonicalDomain The domain that the exchange has declared to be canonical (i.e. what should be used in ads.txt files).
}

// compareCName compare adSystem Canonical Name to the specified domain
func (a adSystem) compareCName(domain string) bool {
	lcDomain := strings.ToLower(domain)

	var match bool
	cNames := strings.Split(a.CanonicalDomain, ",")
	for _, cName := range cNames {
		if strings.TrimSpace(cName) == lcDomain {
			match = true
			break
		}
	}
	return match
}

// newAdSystem init new AdSystem record
func newAdSystem(id int, name string, canonicalDomain string) *adSystem {
	return &adSystem{ID: id, Name: name, CanonicalDomain: canonicalDomain}
}

// adSystemDomain holds known canonical names for AdSystem records
type adSystemDomain struct {
	Domain string // Domain holds the canonical name of a AdSystem records
	ID     int    // ID store the Id of AdSystem record with canonical name equals to this record Name
}

// newAdSystemDomain init new AdSystemDomain record
func newAdSystemDomain(domain string, id int) *adSystemDomain {
	return &adSystemDomain{Domain: domain, ID: id}
}

// ValidateDomainName validates that specified domain name is valid
func validateDomainName(domain string) bool {
	// validate domain has no schema (http(s):\\)
	if strings.Index(domain, "://") != -1 {
		return false
	}

	// parse domain
	u, err := url.Parse("http://" + domain)
	if err != nil {
		log.Printf("Error when parsing [%s] [%s]", domain, err.Error())
		return false
	}

	// make sure that parsed URL host is equal to domain name
	return u.Host == domain
}

// RootDomain Extract “root domain” from specified URL. Root domain is defined as the “public suffix” plus one sting in the name.
func rootDomain(rawurl string) (string, error) {
	// Strip domain from specified URL: remove HTTP schema (http/s) and path from input URL string
	stripDomain := func(rawurl string) string {
		var index int
		// remove schema
		index = strings.Index(rawurl, "://")
		if index != -1 {
			rawurl = rawurl[index+3:]
		}

		// remove path
		index = strings.Index(rawurl, "/")
		if index != -1 {
			rawurl = rawurl[0:index]
		}

		// remove port
		index = strings.Index(rawurl, ":")
		if index != -1 {
			rawurl = rawurl[0:index]
		}

		return rawurl
	}

	// Some private eTLD's are used, eg Amazons. Below should avoid these throwing an error
	_, err := publicsuffix.EffectiveTLDPlusOne(stripDomain(rawurl))
	if err != nil {
		tld, _ := publicsuffix.PublicSuffix(stripDomain(rawurl))
		return tld, nil
	}
	// extract top level domain
	return publicsuffix.EffectiveTLDPlusOne(stripDomain(rawurl))
}

// VaidateAdSystemCName validate that the specifiied ad system domain is a known Ad System.
// It does not imply that any of the ad systems have been vetted or certified.
func vaidateAdSystemCName(domain string) error {
	// check case insensative for domain in
	adSystemDomain, ok := adSystemDomains[strings.ToLower(domain)]
	if !ok {
		// if domain name not found in ad system domains collection, search for it directly in the AdSystem list
		for _, adSystem := range adSystems {
			if adSystem.compareCName(domain) {
				return nil
			}
		}
		return fmt.Errorf("Please verify that %s is a known exchange domain", domain)
	}

	// find domain name in the list of known ad systems
	adSystem, ok := adSystems[adSystemDomain.ID]
	if !ok {
		return fmt.Errorf("Please verify that %s is a known exchange domain", domain)
	}

	// domain does not match Ad System Canonical name: it is still valid but publisher should probably use canonical name
	if len(adSystem.CanonicalDomain) > 0 {
		match := adSystem.compareCName(domain)

		if !match {
			return fmt.Errorf("%s is not the preferred form of the exchange domain. Please consider using %s as the canonical domain name",
				domain, adSystem.CanonicalDomain)
		}
	}

	return nil
}
