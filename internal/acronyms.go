package internal

import (
	"errors"
)

type Acronym struct {
	Key  string
	Word string
}

var acronyms = []Acronym{
	{Key: "lol", Word: "Laugh out loud"},
	{Key: "idk", Word: "I don't know"},
	{Key: "smh", Word: "Shaking my head"},
	{Key: "tbh", Word: "To be honest"},
	{Key: "idc", Word: "I don't care"},
	{Key: "ttyl", Word: "Talk to you later"},
	{Key: "brb", Word: "Be right back"},
	{Key: "omg", Word: "Oh my god"},
	{Key: "wtf", Word: "What the f"},
	{Key: "wth", Word: "What the hell"},
	{Key: "ily", Word: "I love you"},
	{Key: "imy", Word: "I miss you"},
	{Key: "irl", Word: "In real life"},
	{Key: "imo", Word: "In my opinion"},
	{Key: "imho", Word: "In my humble opinion"},
	{Key: "icymi", Word: "In case you missed it"},
	{Key: "tbt", Word: "Throwback Thursday"},
	{Key: "fomo", Word: "Fear of missing out"},
	{Key: "yolo", Word: "You only live once"},
	{Key: "bff", Word: "Best friends forever"},
	{Key: "btw", Word: "By the way"},
	{Key: "faq", Word: "Frequently asked questions"},
	{Key: "diy", Word: "Do it yourself"},
	{Key: "eta", Word: "Estimated time of arrival"},
	{Key: "tgif", Word: "Thank god it's Friday"},
	{Key: "rofl", Word: "Rolling on the floor laughing"},
	{Key: "stfu", Word: "Shut the f*** up"},
	{Key: "tldr", Word: "Too long, didn’t read"},
	{Key: "tmi", Word: "Too much information"},
	{Key: "afaik", Word: "As far as I know"},
	{Key: "lmk", Word: "Let me know"},
	{Key: "nvm", Word: "Nevermind"},
	{Key: "ftw", Word: "For the win"},
	{Key: "byob", Word: "Bring your own beer"},
	{Key: "bogo", Word: "Buy one get one"},
	{Key: "jw", Word: "Just wondering"},
	{Key: "tbf", Word: "To be frank"},
	{Key: "rn", Word: "Right now"},
	{Key: "fubar", Word: "F***** up beyond all recognition"},
	{Key: "iso", Word: "In search of"},
	{Key: "brt", Word: "Be right there"},
	{Key: "ftfy", Word: "Fixed that for you"},
	{Key: "gg", Word: "Good game"},
	{Key: "bfd", Word: "Big freaking deal"},
	{Key: "dae", Word: "Does anyone else?"},
	{Key: "ngl", Word: "Not gonna lie"},
	{Key: "bts", Word: "Behind the scenes"},
	{Key: "ikr", Word: "I know right"},
	{Key: "hmu", Word: "Hit me up"},
	{Key: "fwiw", Word: "For what it’s worth"},
	{Key: "wyd", Word: "What are you doing?"},
	{Key: "idgaf", Word: "I don’t give a f***"},
	{Key: "nbd", Word: "No big deal"},
	{Key: "tba", Word: "To be announced"},
	{Key: "tbd", Word: "To be decided"},
	{Key: "afk", Word: "Away from keyboard"},
	{Key: "abt", Word: "About"},
	{Key: "iykyk", Word: "If you know you know"},
	{Key: "b4", Word: "Before"},
	{Key: "bc", Word: "Because"},
	{Key: "jic", Word: "Just in case"},
	{Key: "snafu", Word: "Situation normal, all f***** up"},
	{Key: "gtg", Word: "Got to go"},
	{Key: "h8", Word: "Hate"},
	{Key: "lmao", Word: "Laughing my a** off"},
	{Key: "iykwim", Word: "If you know what I mean"},
	{Key: "myob", Word: "Mind your own business"},
	{Key: "pov", Word: "Point of view"},
	{Key: "tlc", Word: "Tender loving care"},
	{Key: "hbd", Word: "Happy birthday"},
	{Key: "w/e", Word: "Whatever"},
	{Key: "wysiwyg", Word: "What you see is what you get"},
	{Key: "fwif", Word: "For what it’s worth"},
	{Key: "tw", Word: "Trigger warning"},
	{Key: "eod", Word: "End of day"},
	{Key: "aka", Word: "Also known as"},
	{Key: "asap", Word: "As soon as possible"},
	{Key: "lmgtfy", Word: "Let me Google that for you"},
	{Key: "np", Word: "No problem"},
	{Key: "n/a", Word: "Not applicable or not available"},
	{Key: "ooo", Word: "Out of office"},
	{Key: "tia", Word: "Thanks in advance"},
	{Key: "cob", Word: "Close of business"},
	{Key: "fyi", Word: "For your information"},
	{Key: "nsfw", Word: "Not safe for work"},
	{Key: "wfh", Word: "Work from home"},
	{Key: "omw", Word: "On my way"},
	{Key: "wdyt", Word: "What do you think?"},
	{Key: "wygam", Word: "When you get a minute"},
	{Key: "smp", Word: "Social media platform"},
	{Key: "dm", Word: "Direct message"},
	{Key: "fb", Word: "Facebook"},
	{Key: "ig", Word: "Instagram"},
	{Key: "li", Word: "LinkedIn"},
	{Key: "yt", Word: "YouTube"},
	{Key: "sc", Word: "Snapchat"},
	{Key: "br", Word: "BeReal"},
	{Key: "wa", Word: "WhatsApp"},
	{Key: "tt", Word: "TikTok"},
	{Key: "ff", Word: "Follow Friday"},
	{Key: "im", Word: "Instant message"},
	{Key: "pm", Word: "Private message"},
	{Key: "op", Word: "Original post"},
	{Key: "qotd", Word: "Quote of the day"},
	{Key: "ootd", Word: "Outfit of the day"},
	{Key: "rt", Word: "Retweet"},
	{Key: "tbt", Word: "Throwback Thursday"},
	{Key: "til", Word: "Today I learned"},
	{Key: "ama", Word: "Ask me anything"},
	{Key: "eli5", Word: "Explain like I’m 5"},
	{Key: "fbf", Word: "Flashback Friday"},
	{Key: "mfw", Word: "My feeling when"},
	{Key: "hmu", Word: "Hit me up"},
	{Key: "grwm", Word: "Get ready with me"},
	{Key: "ily", Word: "I love you"},
	{Key: "mcm", Word: "Man crush Monday"},
	{Key: "wcw", Word: "Woman crush Wednesday"},
	{Key: "bf", Word: "Boyfriend"},
	{Key: "gf", Word: "Girlfriend"},
	{Key: "bae", Word: "Before anyone else"},
	{Key: "lysm", Word: "Love you so much"},
	{Key: "pda", Word: "Public display of affection"},
	{Key: "ltr", Word: "Longterm relationship"},
	{Key: "dtr", Word: "Define the relationship"},
	{Key: "ldr", Word: "Long-distance relationship"},
	{Key: "xoxo", Word: "Hugs and kisses"},
	{Key: "otp", Word: "One true pairing"},
	{Key: "loml", Word: "Love of my life"},
	{Key: "cta", Word: "Call to action"},
	{Key: "ugc", Word: "User-generated content"},
	{Key: "ux", Word: "User experience"},
	{Key: "sms", Word: "Short message service"},
	{Key: "mms", Word: "Multimedia messaging service"},
	{Key: "rcs", Word: "Rich communication services"},
	{Key: "roi", Word: "Return on investment"},
	{Key: "ctr", Word: "Click-through rate"},
	{Key: "cpc", Word: "Cost per click"},
	{Key: "cr", Word: "Conversion rate"},
	{Key: "smb", Word: "Small/medium business"},
	{Key: "tos", Word: "Terms of service"},
	{Key: "scn", Word: "Short Code Number"},
	{Key: "5g", Word: "5th generation"},
	{Key: "tcpa", Word: "Telephone Consumer Protection Act"},
}

func GetAcronyms() []string {
	var words []string
	for _, a := range acronyms {
		words = append(words, a.Key)
	}

	return words
}

func GetAcronymValue(key string) string {
	for _, a := range acronyms {
		if a.Key == key {
			return a.Word
		}
	}

	return ""
}

func Exact(str string) (string, error) {
	for _, a := range acronyms {
		if a.Key == str {
			return a.Word, nil
		}
	}

	return "", errors.New("no exact match found")
}
