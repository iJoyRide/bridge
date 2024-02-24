// Copyright 2024 marcu
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package entities

type cardID string
var NameAndUniqueID = map[string]string{
	"Clubs_2" : "AgADAgADK6_WFw",
	"Clubs_3" : "AgADAwADK6_WFw",
	"Clubs_4" : "AgADBAADK6_WFw",
	"Clubs_5" : "AgADBQADK6_WFw",
	"Clubs_6" : "AgADBgADK6_WFw",
	"Clubs_7" : "AgADBwADK6_WFw",
	"Clubs_8" : "AgADCAADK6_WFw",
	"Clubs_9" : "AgADCQADK6_WFw",
	"Clubs_10" : "AgADCgADK6_WFw",
	"Clubs_11" : "AgADCwADK6_WFw",
	"Clubs_12" : "AgADDAADK6_WFw",
	"Clubs_13"  : "AgADDQADK6_WFw",
	"Clubs_14"  : "AgADAQADK6_WFw",
	"Diamonds_2"  : "AgADKQADK6_WFw",
	"Diamonds_3"  : "AgADKgADK6_WFw",
	"Diamonds_4"  : "AgADKwADK6_WFw",
	"Diamonds_5"  : "AgADLAADK6_WFw",
	"Diamonds_6"  : "AgADLQADK6_WFw",
	"Diamonds_7"  : "AgADLgADK6_WFw",
	"Diamonds_8"  : "AgADLwADK6_WFw",
	"Diamonds_9"  : "AgADMAADK6_WFw",
	"Diamonds_10"  : "AgADMQADK6_WFw",
	"Diamonds_11"  : "AgADMgADK6_WFw",
	"Diamonds_12"  : "AgADMwADK6_WFw",
	"Diamonds_13"  : "AgADNAADK6_WFw",
	"Diamonds_14"  : "AgADKAADK6_WFw",
	"Hearts_2"  : "AgADDwADK6_WFw",
	"Hearts_3"  : "AgADEAADK6_WFw",
	"Hearts_4"  : "AgADEQADK6_WFw",
	"Hearts_5"  : "AgADEgADK6_WFw",
	"Hearts_6"  : "AgADEwADK6_WFw",
	"Hearts_7"  : "AgADFAADK6_WFw",
	"Hearts_8"  : "AgADFQADK6_WFw",
	"Hearts_9"  : "AgADFgADK6_WFw",
	"Hearts_10"  : "AgADFwADK6_WFw",
	"Hearts_11"  : "AgADGAADK6_WFw",
	"Hearts_12"  : "AgADGQADK6_WFw",
	"Hearts_13"  : "AgADGgADK6_WFw",
	"Hearts_14"  : "AgADDgADK6_WFw",
	"Spades_2"  : "AgADHAADK6_WFw",
	"Spades_3"  : "AgADHQADK6_WFw",
	"Spades_4"  : "AgADHgADK6_WFw",
	"Spades_5"  : "AgADHwADK6_WFw",
	"Spades_6"  : "AgADIAADK6_WFw",
	"Spades_7"  : "AgADIQADK6_WFw",
	"Spades_8"  : "AgADIgADK6_WFw",
	"Spades_9"  : "AgADIwADK6_WFw",
	"Spades_10"  : "AgADJAADK6_WFw",
	"Spades_11"  : "AgADJQADK6_WFw",
	"Spades_12"  : "AgADJgADK6_WFw",
	"Spades_13"  : "AgADJwADK6_WFw",
	"Spades_14"  : "AgADGwADK6_WFw",
}

var NameAndID = map[string]string {
	"Clubs_2" : "CAACAgIAAxkBAAIEx2XVq4Kxg05Hmifodn3IzfV1M4QUAAICAAMrr9YXwR9KhgSENds0BA",
	"Clubs_3" : "CAACAgIAAxkBAAIEyWXVq-uuRXMOl1eXAYZIMDtRuxADAAIDAAMrr9YXCZeW78C9rOM0BA",
	"Clubs_4" : "CAACAgIAAxkBAAIEymXVrAQN63tBkvvoUpkGuZ2OCcqJAAIEAAMrr9YXVjhLGxQOqt40BA",
	"Clubs_5" : "CAACAgIAAxkBAAIEy2XVrDGGe-llkn9t8x8CD-4p6Fj_AAIFAAMrr9YXI6W4lvOhY6g0BA",
	"Clubs_6" : "CAACAgIAAxkBAAIEzGXVrDixQKZRebLDD70hLGGmzhKTAAIGAAMrr9YXjCzV_E6bATs0BA",
	"Clubs_7" : "CAACAgIAAxkBAAIEzWXVrEOVx7Noz_wVd9oifd5F3okpAAIHAAMrr9YXT7KD1O-4AUc0BA",
	"Clubs_8" : "CAACAgIAAxkBAAIEzmXVrEoblPeeCqxxrv5_JCHilPeRAAIIAAMrr9YXOwiOxdrhWnM0BA",
	"Clubs_9" : "CAACAgIAAxkBAAIEz2XVrFMQEdMVBfhbfvv7P2VC0_CWAAIJAAMrr9YX40LZnq1ecZc0BA",
	"Clubs_10" : "CAACAgIAAxkBAAIE0GXVrFr-SoUUyBAvoS69gdawMB9tAAIKAAMrr9YXTC8vNi1OhT40BA",
	"Clubs_11" : "CAACAgIAAxkBAAIE0WXVrGEb2uAUndsV8qzUmSLH4kXgAAILAAMrr9YX-aVzSWVmLGY0BA",
	"Clubs_12" : "CAACAgIAAxkBAAIE0mXVrGg_Qc6Hl3X_7gu8QL1h2lKmAAIMAAMrr9YXeWEtgUTZnVc0BA",
	"Clubs_13"  : "CAACAgIAAxkBAAIE02XVrG4GybEoYMINj9Zfp7jRySLoAAINAAMrr9YXG2rT4Mt1kAABNAQ",
	"Clubs_14"  : "CAACAgIAAxkBAAIFkGXVve0Gozcn2S2qmqZ8gi3he8okAAIBAAMrr9YX9Zs3aG0XzZ40BA",
	"Diamonds_2"  : "CAACAgIAAxkBAAIE1WXVrPxORkSXCML9OcbmdmL_UTdeAAIpAAMrr9YX7UWi_MPSU8Q0BA",
	"Diamonds_3"  : "CAACAgIAAxkBAAIE1mXVrQPHxuxZDUFG-ZDgLolsf4SVAAIqAAMrr9YXkqPuxsq1OCk0BA",
	"Diamonds_4"  : "CAACAgIAAxkBAAIE12XVrQqp1UGuVM2IBlCDrTwyih6GAAIrAAMrr9YXtg_Y5UXUEkI0BA",
	"Diamonds_5"  : "CAACAgIAAxkBAAIE2GXVrRKEWNE9cTlC1cPV20KHhZ9kAAIsAAMrr9YXNEQuNaZ3Uw00BA",
	"Diamonds_6"  : "CAACAgIAAxkBAAIE2WXVrRrKJoXEKB8hLUux3Sxs9uS0AAItAAMrr9YXu0nGPeN39EU0BA",
	"Diamonds_7"  : "CAACAgIAAxkBAAIE2mXVrSBR-BC3Mpaa-lNE_ph8DzRJAAIuAAMrr9YXFXqMqQEB6AM0BA",
	"Diamonds_8"  : "CAACAgIAAxkBAAIE22XVrSjNC0PXwtfle2XWnOWMWbifAAIvAAMrr9YXB6xnk25-GFU0BA",
	"Diamonds_9"  : "CAACAgIAAxkBAAIE3GXVrS6ND6zaad9tZyr8Nq-j27GoAAIwAAMrr9YXgXnidgr55oI0BA",
	"Diamonds_10"  : "CAACAgIAAxkBAAIE3WXVrTSfXhlWEtH-iss93S6QEGvoAAIxAAMrr9YX0LkZr320yFA0BA",
	"Diamonds_11"  : "CAACAgIAAxkBAAIE3mXVrToWq-75d3Un27t-ikYBx0coAAIyAAMrr9YXGAfhHw5e_bM0BA",
	"Diamonds_12"  : "CAACAgIAAxkBAAIE32XVrUFd_8I9Ief8oY-k0FVUUvSxAAIzAAMrr9YX8-wR7muT2V00BA",
	"Diamonds_13"  : "CAACAgIAAxkBAAIE4GXVrUe9qD6yZqIXhVBnAR8Dl_AEAAI0AAMrr9YXK00p2-xi46s0BA",
	"Diamonds_14"  : "CAACAgIAAxkBAAIE4WXVrU_J3RAmoj-wOHOrRQFEC7PcAAIoAAMrr9YXzzQAAWR1ImN_NAQ",
	"Hearts_2"  : "CAACAgIAAxkBAAIE4mXVrVYRENez1nwheEmqQH9yKFQRAAIPAAMrr9YXH3Ydru8YTnU0BA",
	"Hearts_3"  : "CAACAgIAAxkBAAIE42XVrV79lux197NGShWF_-b8xIIGAAIQAAMrr9YXWEzmIjTZIZE0BA",
	"Hearts_4"  : "CAACAgIAAxkBAAIE5GXVrWQ9Ibvi__pL0pGEyHWoOv2-AAIRAAMrr9YXfTSUWXvsb_k0BA",
	"Hearts_5"  : "CAACAgIAAxkBAAIE5WXVrXDqhfMwzd6zhdQLnCJEE-5RAAISAAMrr9YXlcPcr7UnbIY0BA",
	"Hearts_6"  : "CAACAgIAAxkBAAIE5mXVrXazPSdFho0W81bd8JIIS1J7AAITAAMrr9YX9S9KB7eRRA00BA",
	"Hearts_7"  : "CAACAgIAAxkBAAIE52XVrX3XGdIpKgK_zc1zpnraMahKAAIUAAMrr9YX0Y-0RBpJTLQ0BA",
	"Hearts_8"  : "CAACAgIAAxkBAAIE6GXVrYSYio8qDRqt8pndKXesEARFAAIVAAMrr9YXb1dmdl8O-G00BA",
	"Hearts_9"  : "CAACAgIAAxkBAAIE6WXVrYqzxZa2PVTXIJyTudGaGIOrAAIWAAMrr9YX77CLDfMNMLo0BA",
	"Hearts_10"  : "CAACAgIAAxkBAAIE6mXVrZAmiNoehuvOQou21iMDU1RpAAIXAAMrr9YXzUvWVFLbQow0BA",
	"Hearts_11"  : "CAACAgIAAxkBAAIE62XVrZVO0iPJr1i2gSF2WVFstR3LAAIYAAMrr9YXPZCIBug_nfM0BA",
	"Hearts_12"  : "CAACAgIAAxkBAAIE7GXVrZyff9rX6WpcvqWIBQmT8vZ4AAIZAAMrr9YXStQidqI_7wU0BA",
	"Hearts_13"  : "CAACAgIAAxkBAAIE7WXVraRO1V1XdbvAEwXQrCQf6n5oAAIaAAMrr9YXAviFEp8-Fm80BA",
	"Hearts_14"  : "CAACAgIAAxkBAAIE7mXVrar3Yamd-fL9IGF0I2czvhK3AAIOAAMrr9YXfvu9md3kZrw0BA",
	"Spades_2"  : "CAACAgIAAxkBAAIE72XVrbrDq_jHEFGNjAv6ybSiVIf-AAIcAAMrr9YXXrafHp0orAk0BA",
	"Spades_3"  : "CAACAgIAAxkBAAIE8GXVrcEyBxhCWhu3J4RSbAzArm8_AAIdAAMrr9YXHGaTmsvkKBU0BA",
	"Spades_4"  : "CAACAgIAAxkBAAIE8WXVrcgf7UQ9MVvrTViJmeXiPEpCAAIeAAMrr9YXRkTIXSbUfrg0BA",
	"Spades_5"  : "CAACAgIAAxkBAAIE8mXVrc8Sbon4msUT2tp7ba67cwdpAAIfAAMrr9YXb3Q1IqDIiQE0BA",
	"Spades_6"  : "CAACAgIAAxkBAAIE82XVrdXYnR2vWw32maKRdGht0G0OAAIgAAMrr9YX0CEXYmgUKWE0BA",
	"Spades_7"  : "CAACAgIAAxkBAAIE9GXVrdzmUoE9Zk-yErDd2SzLRPg4AAIhAAMrr9YXoepgqQ6muAM0BA",
	"Spades_8"  : "CAACAgIAAxkBAAIE9WXVrePTZDoJF6qXp8d5hydsEX0WAAIiAAMrr9YXuXmqmH6w7ns0BA",
	"Spades_9"  : "CAACAgIAAxkBAAIE9mXVrekdZHq7n35Fy9zn4oUHFjyWAAIjAAMrr9YXVb1756Abx6o0BA",
	"Spades_10"  : "CAACAgIAAxkBAAIE92XVrfAuZfrfbiSdE_jPmpS47kvpAAIkAAMrr9YX9f6AXNbvQ_A0BA",
	"Spades_11"  : "CAACAgIAAxkBAAIE-GXVrfe4Kb6f7Cx1JYJCVUXfawABcQACJQADK6_WF4PZPLZKQFt4NAQ",
	"Spades_12"  : "CAACAgIAAxkBAAIE-mXVrgABs1DPCgABdbRUh2LTP7jmxOMAAiYAAyuv1hfCfweOI6wFGzQE",
	"Spades_13"  : "CAACAgIAAxkBAAIE-2XVrga7KTOH9J508wvf2bzkiImlAAInAAMrr9YXHfZjvnv9-UM0BA",
	"Spades_14"  : "CAACAgIAAxkBAAIE_GXVrg0cqmKigUSq9PnN20Rbhv1BAAIbAAMrr9YX1TA9KcRoFD40BA",
}


func NameToID(card string) string {
	for key,value := range NameAndID{
		if key == card{
			return value
		}
	}
	return ""
}

func IDToName(cardID string) string {
	for key,value := range NameAndUniqueID{
		if value == cardID{
			return key
		}
	}
	return ""
}

