package models

import (
	"time"

	gormigrate "github.com/genofire/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	// stream

	testdataStream1      = uuid.MustParse("dffe2c0e-3713-4399-8ee2-279becbbb06e")
	testdataStream1Lang1 = uuid.MustParse("3a4f9157-65bf-4d15-a82b-1cd9295d07e0")

	testdataStream2 = uuid.MustParse("0801a547-59f1-4a63-946f-2ab03f62e6ee")

	testdataStream3         = uuid.MustParse("06e3a71e-581d-4735-9647-3e4a49b5caa8")
	testdataStream3Lang1    = uuid.MustParse("cff00fcd-5408-4cb4-8ac7-2d42b45fbd68")
	testdataStream3Speaker1 = uuid.MustParse("c03aa102-551e-4b3a-b670-5b6c7ac13faa")

	testdataStream4         = uuid.MustParse("57de7dfd-c060-4da1-8f57-f0880c1f2e5e")
	testdataStream4Lang1    = uuid.MustParse("d92fbc3b-a027-49f6-849b-7efb425aa5c0")
	testdataStream4Speaker1 = uuid.MustParse("0d1b38cd-561c-4db4-b4b9-51f74ba3dba4")
	testdataStream4Speaker2 = uuid.MustParse("1dbf0438-a9c1-4412-b44c-08fe7819902c")

	testdataStream5      = uuid.MustParse("4fb029d6-063a-4302-9ae8-4c1c6a1542a5")
	testdataStream5Lang1 = uuid.MustParse("d5262bb7-378b-456f-9e91-34f63b174c48")

	testdataStream6      = uuid.MustParse("1742d9b6-c9c6-45fb-a3a3-4a3e7fac2987")
	testdataStream6Lang1 = uuid.MustParse("0b7136a6-4c51-49ac-99e9-27ef833169f6")

	// Grand Piano: Martin Kohlstedt
	testdataStream7         = uuid.MustParse("7fc21416-5d68-4ecf-bd4b-e8a89f7c31f2")
	testdataStream7Lang1    = uuid.MustParse("45bb0727-4812-40a2-8ac2-12422dfb42f1")
	testdataStream7Speaker1 = uuid.MustParse("976010a0-c19f-4d22-a4d6-9553b460adfe")

	// Grand Piano: Lucia Cadotsch
	testdataStream8         = uuid.MustParse("7ff58740-8c3a-4e09-8fc1-1eeb39c2a9d4")
	testdataStream8Lang1    = uuid.MustParse("b7be68f2-a109-4e28-8744-bfc6c8f03f9f")
	testdataStream8Speaker1 = uuid.MustParse("dfa0ff16-8cb0-46e2-a56a-e44dcde1868e")
	testdataStream8Speaker2 = uuid.MustParse("bd72ee71-6e0f-4ba6-9d3d-eb6b7ba589e3")
	testdataStream8Speaker3 = uuid.MustParse("96c157f1-10d3-4ae1-af59-81ccece5d4fc")

	// Grand Piano: Pablo Ortega
	testdataStream9         = uuid.MustParse("36ba6bfe-2b40-425d-8cc7-d7de5ec4b67a")
	testdataStream9Lang1    = uuid.MustParse("c6720db6-9d62-483d-a6ec-5cd1137b4dac")
	testdataStream9Speaker1 = uuid.MustParse("e7e6eb6b-8188-4169-ae4e-8144d559a592")

	// Grand Piano: Pulsar Trio
	testdataStream10         = uuid.MustParse("f4d8d0e2-47bf-412c-a5e1-80c0f85a9e4c")
	testdataStream10Lang1    = uuid.MustParse("619e61ea-a6de-42e3-b48f-6b5e52f8d920")
	testdataStream10Speaker1 = uuid.MustParse("09120c00-acb6-4865-b1b5-6bf39963e47f")

	// Grand Piano: Motschmann Trio
	testdataStream11         = uuid.MustParse("710b445a-51f5-4a9c-8fd4-59956453401c")
	testdataStream11Lang1    = uuid.MustParse("6c0c2394-bc5c-43b0-ae14-8957118f8231")
	testdataStream11Speaker1 = uuid.MustParse("54e80294-b92d-4dfe-9c67-23e1824b8e4f")

	// Grand Piano: Niklas Paschburg
	testdataStream12         = uuid.MustParse("54ff055b-5e46-4344-a43f-deb41c693045")
	testdataStream12Lang1    = uuid.MustParse("7b4275e0-2a11-4588-8264-51ff699d5868")
	testdataStream12Speaker1 = uuid.MustParse("6bd38420-f647-42ed-ba7c-99bd17b3cfe7")

	/* Grand Piano: -
	testdataStream9         = uuid.MustParse("")
	testdataStream9Lang1    = uuid.MustParse("")
	testdataStream9Speaker1 = uuid.MustParse("")
	*/
)

func init() {

	testdata = append(testdata, []*gormigrate.Migration{
		{
			ID: "10-data-0030-01-stream-1",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataStream1,
					ChannelID: TestChannelID1,
					Chat:      false,
					Running:   true,
					StartAt:   time.Date(2021, 3, 4, 18, 30, 0, 0, loc),
					ListenAt:  time.Date(2021, 3, 4, 18, 30, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/728edaf7-9ad9-f972-4d09-ba5940cd43f9/poster.png",
					Preview:   "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/728edaf7-9ad9-f972-4d09-ba5940cd43f9/preview.webp",
					Tags:      []*Tag{{ID: TestTagVortragID}},
					Speakers: []*Speaker{
						{
							ID: testdataRecording7Speaker1,
						},
						{
							ID: testdataRecording7Speaker2,
						},
						{
							ID: testdataRecording7Speaker3,
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataStream1Lang1,
					StreamID: testdataStream1,
					Lang:     "de",
					Title:    "„Die mir von der Wehrmacht angebotenen Kriegsgefangenen sind derart entkräftet“",
					Subtitle: "Sowjetische Kriegsgefangene in Bremer Arbeitskommandos 1941-1945",
					Short:    `Sowjetische Kriegsgefangene bildeten eine der größten Opfergruppen des Nationalsozialismus. Die Wehrmacht brachte Millionen sowjetische Soldat\*innen zum Arbeitseinsatz ins Deutsche Reich. Mehr als die Hälfte von ihnen überlebte die Kriegsgefangenschaft in den Kriegsgefangenenlagern wie dem Stalag X B Sandbostel und den Außenkommandos nicht. Auch in Bremen setzten Firmen und Behörden die kriegsgefangenen Rotarmisten zur Arbeit ein, vornehmlich in der Rüstungsindustrie. Im unserem Vortrag wollen wir die ökonomischen und ideologischen Hintergründe und Widersprüche dieser Arbeitseinsätze aufzeigen. ...`,
					Long: `Sowjetische Kriegsgefangene bildeten eine der größten Opfergruppen des Nationalsozialismus. Die Wehrmacht brachte Millionen sowjetische Soldat\*innen zum Arbeitseinsatz ins Deutsche Reich. Mehr als die Hälfte von ihnen überlebte die Kriegsgefangenschaft in den Kriegsgefangenenlagern wie dem Stalag X B Sandbostel und den Außenkommandos nicht. Auch in Bremen setzten Firmen und Behörden die kriegsgefangenen Rotarmisten zur Arbeit ein, vornehmlich in der Rüstungsindustrie. Im unserem Vortrag wollen wir die ökonomischen und ideologischen Hintergründe und Widersprüche dieser Arbeitseinsätze aufzeigen. Anhand einzelner exemplarischer Arbeitskommandos beleuchten wir die Lebens- und Arbeitsbedingungen von sowjetischen Kriegsgefangenen in Bremen. Der Vortrag lädt alle Interessierte zum Austausch über dieses lange verdrängte Thema ein.

Eine Veranstaltung der Gedenkstätte Lager Sandbostel in Kooperation mit dem Kulturzentrum Kukoon.

Bildinfo: Personalkarte des sowjetischen Kriegsgefangenen Wasilij M. Alexejew, der am 15.09.1942 in das Arbeitskommando der Bremer Francke-Werke eingesetzt wurde und am 11.03.1942 an Tuberkulose starb, Quelle [https://obd-memorial.ru/html/info.htm?id=300643349](https://obd-memorial.ru/html/info.htm?id=300643349)`,
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Delete(&StreamLang{
					ID: testdataStream1Lang1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Exec("DELETE FROM stream_speakers WHERE stream_id = ?", testdataStream1).Error; err != nil {
					return err
				}
				if err := tx.Exec("DELETE FROM stream_tags WHERE stream_id = ?", testdataStream1).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Stream{
					ID: testdataStream1,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "10-data-0030-01-stream-2",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataStream2,
					ChannelID: TestChannelID2,
					Chat:      false,
					Running:   true,
					StartAt:   time.Date(2021, 3, 4, 18, 30, 0, 0, loc),
					ListenAt:  time.Date(2021, 3, 4, 18, 30, 0, 0, loc),
					Poster:    "https://c3woc.de/images/banner.jpg",
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Exec("DELETE FROM stream_speakers WHERE stream_id = ?", testdataStream2).Error; err != nil {
					return err
				}
				if err := tx.Exec("DELETE FROM stream_tags WHERE stream_id = ?", testdataStream2).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Stream{
					ID: testdataStream2,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "10-data-0030-01-stream-3",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataStream3,
					ChannelID: TestChannelID1,
					Chat:      true,
					Running:   true,
					StartAt:   time.Date(2021, 4, 15, 0, 0, 0, 0, loc),
					ListenAt:  time.Date(2021, 4, 15, 0, 0, 0, 0, loc),
					Poster:    "https://media.kukoon.de/images/41185fa4-3e22-44bb-9020-1d824e12ede3.jpg",
					Tags: []*Tag{
						{ID: TestTagBuchvorstellungID},
						{ID: TestTagDiskussionID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataStream3Speaker1,
							Name:    "Andreas Speit",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataStream3Lang1,
					StreamID: testdataStream3,
					Lang:     "de",
					Title:    "Rechte Egoshooter",
					Subtitle: "Von der virtuellen Hetze zum Livestream-Attentat",
					Short:    `Weltweit gibt es rechtsterroristische Attentate eines neuen Typs. In Halle (Saale) verhinderte nur eine verschlossene Holztür der Synagoge ein größeres Massaker. ...`,
					Long: `Weltweit gibt es rechtsterroristische Attentate eines neuen Typs. In Halle (Saale) verhinderte nur eine verschlossene Holztür der Synagoge ein größeres Massaker. Am 9. Oktober 2019 wollte dort ein Rechtsextremist die versammelten Juden hinrichten. Mit selbstgebauten Waffen schoss er auf die Tür und warf eigens hergestellte Sprengsätze. Online konnten Gleichgesinnte zusehen, wie er zwei Menschen ermordete: Seine Tat verbreitete er per Videokamera auf einem Portal für Computerspiel-Videos. Er ahmte damit andere »Egoshooter« nach – wie einen Rechtsextremisten, der in Neuseeland wenige Monate zuvor die Tötung von 51 Menschen live im Internet übertragen hatte. Was treibt Menschen vom Bildschirm zur realen Gewalt auf der Straße? Die Beiträge des Buches gehen den Spuren der Attentäter nach und zeigen die speziellen Radikalisierungsmechanismen im Netz auf. Sie erklären die Hintergründe und Motive dieser Männer, die in ihren rechten Online-Gemeinden Antisemitismus, Rassismus und Antifeminismus verbreiten. Das Buch gibt Einblicke in eine Welt, die vielen unbekannt ist.

Eine Veranstaltung des _Kulturzentrum Kukoon_ in Kooperation mit der _Rosa-Luxemburg-Initiative – Die Rosa-Luxemburg-Stiftung in Bremen_.
				`,
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Delete(&StreamLang{
					ID: testdataStream3Lang1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Exec("DELETE FROM stream_speakers WHERE stream_id = ?", testdataStream3).Error; err != nil {
					return err
				}
				if err := tx.Exec("DELETE FROM stream_tags WHERE stream_id = ?", testdataStream3).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Stream{
					ID: testdataStream3,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "10-data-0030-01-stream-4",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataStream4,
					ChannelID: TestChannelID1,
					Chat:      true,
					Running:   true,
					StartAt:   time.Date(2021, 4, 22, 0, 0, 0, 0, loc),
					ListenAt:  time.Date(2021, 4, 22, 0, 0, 0, 0, loc),
					Poster:    "https://media.kukoon.de/images/f61fa2de-72d3-4a1e-98b2-65b13d8ecb01.jpg",
					Tags: []*Tag{
						{ID: TestTagBuchvorstellungID},
						{ID: TestTagDiskussionID},
					},
					Speakers: []*Speaker{
						{
							ID: testdataStream4Speaker1,
						},
						{
							ID: testdataStream4Speaker2,
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataStream4Lang1,
					StreamID: testdataStream4,
					Lang:     "de",
					Title:    "Aufklären und Einmischen",
					Subtitle: "Der NSU-Komplex und der Münchner Prozess - Buchvorstellung mit NSU-Watch",
					Short:    `Im November 2011 kam eine rechtsterroristische Mord- und Anschlagsserie des sogenannten Nationalsozialistischen Untergrunds (NSU) ans Licht, die in ihrer Dimension neu war. In den folgenden Untersuchungen formte sich ein erstes Bild des NSU-Komplexes. ...`,
					Long: `Im November 2011 kam eine rechtsterroristische Mord- und Anschlagsserie des sogenannten Nationalsozialistischen Untergrunds (NSU) ans Licht, die in ihrer Dimension neu war. In den folgenden Untersuchungen formte sich ein erstes Bild des NSU-Komplexes. Dabei wurde deutlich, dass eine noch umfassendere juristische und gesellschaftliche Aufarbeitung anstand. So beschlossen antifaschistische Initiativen und Einzelpersonen, die Arbeit am NSU-Komplex zu verstetigen, und gründeten »NSU-Watch«. Neun Jahre später ist die Aufarbeitung des NSU-Komplexes noch lange nicht abgeschlossen, die Gefahr des rechten Terrors bleibt schrecklich aktuell. NSU-Watch hat den NSU-Prozess beobachtet, jeden Tag protokolliert und der Öffentlichkeit zur Verfügung gestellt. Darüber hinaus haben sich Landesprojekte gegründet, die die parlamentarischen Aufklärungsbemühungen begleiten. Das zentrale Anliegen des Buches von NSU-Watch ist, die rassistischen Strukturen, die den NSU hervorbrachten, ihn wissentlich oder unwissentlich unterstützten und so zehn Morde, drei Sprengstoffanschläge und 15 Raubüberfälle zwischen 1998 und 2011 möglich machten, entlang der Geschehnisse und Akteur*innen des NSU-Prozesses in München aufzuzeigen. Trotz der vielen offen gebliebenen Fragen soll das Buch eine Zwischenbilanz bieten, die antifaschistischer Demokratieförderung zugrunde gelegt werden kann.

## NSU-Watch
Das Autor\*innen-Kollektiv NSU-WATCH besteht aus Mitgliedern der unabhängigen Beobachtungsstelle NSU-Watch – Aufklären & Einmischen, die sich im Jahr 2012 gegründet hat, um die Aufklärungsbemühungen zum NSU-Komplex zu unterstützen und kritisch zu begleiten. NSU-Watch wird von einem Bündnis aus rund einem Dutzend antifaschistischer und antirassistischer Gruppen und Einzelpersonen aus dem ganzen Bundesgebiet getragen, die teilweise seit Jahrzehnten zum Themenkomplex Rechter Terror arbeiten. Kern der Arbeit von NSU-Watch war bzw. ist die Beobachtung des NSU-Prozesses am Oberlandesgericht in München sowie der diversen parlamentarischen Untersuchungsausschüsse im Bundestag und in den Ländern.

Eine Veranstaltung des _Kulturzentrum Kukoon_ in Kooperation mit der _Rosa-Luxemburg-Initiative – Die Rosa-Luxemburg-Stiftung in Bremen_.
				`,
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Delete(&StreamLang{
					ID: testdataStream4Lang1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Exec("DELETE FROM stream_speakers WHERE stream_id = ?", testdataStream4).Error; err != nil {
					return err
				}
				if err := tx.Exec("DELETE FROM stream_tags WHERE stream_id = ?", testdataStream4).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Stream{
					ID: testdataStream4,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "10-data-0030-01-stream-5",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataStream5,
					ChannelID: TestChannelID1,
					EventID:   &TestEventID1,
					Chat:      true,
					Running:   true,
					StartAt:   time.Date(2021, 5, 5, 0, 0, 0, 0, loc),
					ListenAt:  time.Date(2021, 5, 5, 0, 0, 0, 0, loc),
					Poster:    "https://media.kukoon.de/images/67bd5c4c-81d6-47c8-adb9-458a9da58dbd.jpg",
					Tags: []*Tag{
						{ID: TestTagBuchvorstellungID},
						{ID: TestTagDiskussionID},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataStream5Lang1,
					StreamID: testdataStream5,
					Lang:     "de",
					Title:    "Mareice Kaiser",
					Subtitle: "Das Unwohlsein der modernen Mutter",
					Short:    `Mütter sollen heute alles sein: Versorgerin, Businesswoman, Mom I'd like to fuck. Dass darunter ihr Wohlbefinden leidet, ist kein Wunder. Mareice Kaiser, Journalistin und selbst Mutter, stellt dabei immer wieder fest: ...`,
					Long:     `Mütter sollen heute alles sein: Versorgerin, Businesswoman, Mom I'd like to fuck. Dass darunter ihr Wohlbefinden leidet, ist kein Wunder. Mareice Kaiser, Journalistin und selbst Mutter, stellt dabei immer wieder fest: Das Mutterideal ist unerreichbar und voller Widersprüche. Nichts kann man richtig machen und niemandem etwas recht. Mutterschaft berührt dabei, natürlich, jeden Lebensbereich: Denn egal, ob es um Arbeit, Geld, Sex, Körper, Psyche oder Liebe geht – Stereotype, Klischees und gesellschaftlichen Druck gibt es überall, auf Instagram, im Bett und im Büro. In ihrem Buch "Das Unwohlsein der modernen Mutter" (Rowohlt, 2021) zeigt die Autorin, wo Mütter heute stehen: noch immer öfter am Herd als in den Chefetagen. Und, wo sie stehen sollten: Dort, wo sie selbst sich sehen – frei und selbstbestimmt. Bei OUT LOUD liest Mareice Kaiser aus ihrem Buch und spricht mit uns über Frausein, Mutterschaft und Selbstbestimmung.`,
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Delete(&StreamLang{
					ID: testdataStream5Lang1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Exec("DELETE FROM stream_speakers WHERE stream_id = ?", testdataStream5).Error; err != nil {
					return err
				}
				if err := tx.Exec("DELETE FROM stream_tags WHERE stream_id = ?", testdataStream5).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Stream{
					ID: testdataStream5,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "10-data-0030-01-stream-6",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataStream6,
					ChannelID: TestChannelID1,
					Chat:      true,
					Running:   true,
					StartAt:   time.Date(2021, 6, 24, 0, 0, 0, 0, loc),
					ListenAt:  time.Date(2021, 6, 24, 0, 0, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/1742d9b6-c9c6-45fb-a3a3-4a3e7fac2987/poster.png",
					Tags: []*Tag{
						{ID: TestTagDiskussionID},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataStream6Lang1,
					StreamID: testdataStream6,
					Lang:     "de",
					Title:    "System Change not Climate Change!",
					Subtitle: "Einführung zu Klimakrise und Kapitalismuskritik",
					Short: `Diskussionsveranstaltung mit der Gruppe **direction f** (Hannover)

**Kukoon im Park** oder hier`,
					Long: `Diskussionsveranstaltung mit der Gruppe **direction f** (Hannover)

**Kukoon im Park** oder hier

Zwar verblasst die Klimakrise seit einem Jahr im medialen Schatten der Corona-Pandemie, die Dringlichkeit zum Handeln bleibt jedoch unverändert. Klar ist, dass die Klimakrise kein rein ökologisches Phänomen sondern ebenso sehr eine soziale Krise ist. Als »direction f« haben wir uns bisher vorrangig mit den Zusammenhängen von Klimakrise und Kapitalismus befasst. Im Rahmen der Veranstaltung wollen wir kurz auf den Ist-Zustand und bestehende Zusammenhänge eingehen. Davon ausgehend würden wir gerne darüber diskutieren, was (un)taugliche Strategien gegen die drohende Klimakatstrophe sein können und welche Rolle und Aufgaben dabei einer (radikalen) Linken zukämen. direction f ist ein Zusammenschluss von Menschen in Hannover, der sich bisher schwerpunktmäßig mit dem Zusammenhang von Klimakrise und Kapitalismus befasst hat.

Mehr Infos unter [direction-f.org](https://direction-f.org/)`,
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Delete(&StreamLang{
					ID: testdataStream6Lang1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Exec("DELETE FROM stream_speakers WHERE stream_id = ?", testdataStream6).Error; err != nil {
					return err
				}
				if err := tx.Exec("DELETE FROM stream_tags WHERE stream_id = ?", testdataStream6).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Stream{
					ID: testdataStream6,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "10-data-0030-01-stream-7",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataStream7,
					ChannelID: TestChannelID1,
					EventID:   &TestEventID2,
					Chat:      false,
					Running:   true,
					StartAt:   time.Date(2021, 7, 21, 0, 0, 0, 0, loc),
					ListenAt:  time.Date(2021, 7, 21, 0, 0, 0, 0, loc),
					Poster:    "https://kukoon.de/images/7/6/7/e/5/767e5a264900c670ff18777ee9d5e67466c2a185-martin-kohlstedt.jpg",
					Tags: []*Tag{
						{ID: TestTagKonzertID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataStream7Speaker1,
							Name:    "Martin Kohlstedt",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataStream7Lang1,
					StreamID: testdataStream7,
					Lang:     "de",
					Title:    "Martin Kohlstedt",
					Subtitle: "Mittwoch ab 18:00",
					Short:    `Kein geringerer als **Martin Kohlstedt** wird den wunderbaren Auftakt machen und den Park auf die kommenden Konzerte perfekt einstimmen.`,
					Long: `Kein geringerer als **Martin Kohlstedt** wird den wunderbaren Auftakt machen und den Park auf die kommenden Konzerte perfekt einstimmen.
Das neue Martin Kohlstedt Album »FLUR« erschien im November 2020 auf Warner Classics. Das besondere Setup aus Flügel, Synthesizern und Electronika, kombiniert mit Kohlstedts Ansatz jedes Konzert von Grund auf neu zu verhandeln macht seine Konzerte zu einem Erlebnis.

**Martin Kohlstedt** lebt und arbeitet in Weimar. Seine bisherigen Alben TAG, NACHT, STROM, STRÖME und deren Begleiter in Form von Reworks erhielten internationale Anerkennung und führten den Komponisten und Pianisten auf Konzertreisen in der ganzen Welt.

**Kohlstedt** bezeichnet seine Art des Arbeitens als modulares Komponieren, die Stücke sind ständig in Bewegung und folgen auch im Konzert keiner festgelegten Form. Improvisation ist zwingend Teil des Schaffens des 1988 geborenen Musikers, ebenso wie Augenhöhe mit dem Publikum, der Mut zu Scheitern und die Interaktion mit Raum, Menschen und Kontext.`,
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Delete(&StreamLang{
					ID: testdataStream7Lang1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Exec("DELETE FROM stream_speakers WHERE stream_id = ?", testdataStream7).Error; err != nil {
					return err
				}
				if err := tx.Exec("DELETE FROM stream_tags WHERE stream_id = ?", testdataStream7).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Stream{
					ID: testdataStream7,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "10-data-0030-01-stream-8",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataStream8,
					ChannelID: TestChannelID1,
					EventID:   &TestEventID2,
					Chat:      false,
					Running:   true,
					StartAt:   time.Date(2021, 7, 21, 18, 15, 0, 0, loc),
					ListenAt:  time.Date(2021, 7, 21, 18, 15, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/7ff58740-8c3a-4e09-8fc1-1eeb39c2a9d4/poster.png",
					Tags: []*Tag{
						{ID: TestTagKonzertID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataStream8Speaker1,
							Name:    "Lucia Cadotsch",
						},
						{
							OwnerID: TestChannelID1,
							ID:      testdataStream8Speaker2,
							Name:    "Ronny Graupe",
						},
						{
							OwnerID: TestChannelID1,
							ID:      testdataStream8Speaker3,
							Name:    "Clara Vetter",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataStream8Lang1,
					StreamID: testdataStream8,
					Lang:     "de",
					Title:    "Lucia Cadotsch",
					Subtitle: "Mittwoch ab 20:30",
					Short: `Hierzu ziteren wir den Deutschlandfunk:
__Sie hat dem Jazzgesang neuen Atem eingehaucht und experimentiert ständig weiter.__ (deutschlandfunk.de)`,
					Long: `Am ersten Abend wird uns **Lucia Cadotsch** begleitet von **Ronny Graupe** (Gitarre) in die Nacht begleiten. Dabei werden sie von **Clara Vetter** unterstützt.

Hierzu ziteren wir den Deutschlandfunk:
__Sie hat dem Jazzgesang neuen Atem eingehaucht und experimentiert ständig weiter.__ (deutschlandfunk.de)

1984 wurde **Lucia Cadotsch** in Zürich geboren. Mit 14 entdeckte sie ihre Liebe zum Jazz: Die Plattensammlung ihres Vaters begeisterte sie, vor allem Aufnahmen mit Miles Davis, John Coltrane, Nina Simone und Billie Holiday.
**Lucia Cadotsch** bekam Klavier- und Gesangsunterricht und ging mit 18 Jahren an die Universität der Künste Berlin, um Jazzgesang zu studieren. Sie begründete diverse Ensembles, zum Beispiel das Popquartett Schneeweiss + Rosenrot, mit dem sie 2012 den Neuen Deutschen Jazzpreis gewann.
2016 gelang **Lucia Cadotsch** der internationale Durchbruch mit dem Album »Speak Low«. Ein Jahr später erhielt sie den ECHO Jazz als Sängerin des Jahres. Heute lebt sie in Berlin. `,
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Delete(&StreamLang{
					ID: testdataStream8Lang1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Exec("DELETE FROM stream_speakers WHERE stream_id = ?", testdataStream8).Error; err != nil {
					return err
				}
				if err := tx.Exec("DELETE FROM stream_tags WHERE stream_id = ?", testdataStream8).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Stream{
					ID: testdataStream8,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "10-data-0030-01-stream-9",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataStream9,
					ChannelID: TestChannelID1,
					EventID:   &TestEventID2,
					Chat:      false,
					Running:   true,
					StartAt:   time.Date(2021, 7, 22, 17, 0, 0, 0, loc),
					ListenAt:  time.Date(2021, 7, 22, 17, 0, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/36ba6bfe-2b40-425d-8cc7-d7de5ec4b67a/poster.png",
					Tags: []*Tag{
						{ID: TestTagKonzertID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataStream9Speaker1,
							Name:    "Pablo Ortega",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataStream9Lang1,
					StreamID: testdataStream9,
					Lang:     "de",
					Title:    "Pablo Ortega",
					Subtitle: "Donnerstag ab 18:00",
					Short:    `Pablo Ortega ist ein spanischer Cellist und Komponist. Er lebt in Bremen, wo er aktuell als multidisziplinarischer Musiker aktiv ist.`,
					Long: `Den zweiten Abend wird **Pablo Ortega** eröffnen.

**Pablo Ortega** ist ein spanischer Cellist und Komponist. Er lebt in Bremen, wo er aktuell als multidisziplinarischer Musiker aktiv ist. In seinen Werken verbindet er Elemente moderner klassischer Musik – wie z.B. intime Cellomelodien – mit elektronischen Beats und organischen, atmosphärischen Klangtexturen mit dem Synthesizer. Damit schafft er eine Mischung von Genres, die von filmischer akustischer Musik bis zu energetischer Electronica reicht.

Seine erste EP »Still Waters Run Deep« erschien im Februar 2020. `,
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Delete(&StreamLang{
					ID: testdataStream9Lang1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Exec("DELETE FROM stream_speakers WHERE stream_id = ?", testdataStream9).Error; err != nil {
					return err
				}
				if err := tx.Exec("DELETE FROM stream_tags WHERE stream_id = ?", testdataStream9).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Stream{
					ID: testdataStream9,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "10-data-0030-01-stream-10",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataStream10,
					ChannelID: TestChannelID1,
					EventID:   &TestEventID2,
					Chat:      false,
					Running:   true,
					StartAt:   time.Date(2021, 7, 22, 20, 15, 0, 0, loc),
					ListenAt:  time.Date(2021, 7, 22, 20, 15, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/f4d8d0e2-47bf-412c-a5e1-80c0f85a9e4c/poster.png",
					Tags: []*Tag{
						{ID: TestTagKonzertID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataStream10Speaker1,
							Name:    "Pulsar Trio",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataStream10Lang1,
					StreamID: testdataStream10,
					Lang:     "de",
					Title:    "Pulsar Trio",
					Subtitle: "Donnerstag ab 20:30",
					Short:    `Treibend, pulsierend, vital. Raffinierte Rhythmen, freie Improvisation und Stücke, deren Melodien das Zeug zum Ohrwurm haben `,
					Long:     `Treibend, pulsierend, vital. Raffinierte Rhythmen, freie Improvisation und Stücke, deren Melodien das Zeug zum Ohrwurm haben – mit scheinbar so gegensätzlichen Instrumenten wie Sitar, Piano und Drums lässt das **Pulsar Trio** nicht nur einen neuartigen Klangraum entstehen, sondern auch eine groovige Fusion aus freiem Jazzdenken und originären Worldbeats. »Trotz ihrer vielen Taktwechsel, Breaks und Generalpausen wirkt [die Musik] nie verkopft oder kompliziert, sondern bleibt stets nachvollziehbar und zugänglich, vielfach – dank eines ausgeprägten Gespürs für griffige Melodiebögen – sogar geradezu eingängig.« (Harry Schmidt) Die creole – Preisträger konnten die hohe Qualität ihrer musikalischen Entdeckungsreise auf einer Vielzahl von Konzerten unter Beweis stellen. So gastierten sie u.a. auf den Leverkusener Jazztagen, dem Fusion-Festival, dem Glastonbury Festival, dem Rudolstadt Festival sowie dem Jazzfestival Izmir. Mit ihrem 3. Album »Zoo of Songs«, das im April 2018 bei t3 records erschien, erschließen sich die drei Musiker noch einmal ganz neue musikalische Sphären – kontemplativ, elektronisch, rhythmisch listig, energiegeladen und in satten Klangfarben präsentiert sich das Trio herausragend gereift und mit diesem pointierten Freigeist, der es so unverwechselbar macht. Die »Zoo of Songs« – Record-Release-Tour lässt sie ab 2018 auf internationalen Konzertbühnen unterwegs sein.`,
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Delete(&StreamLang{
					ID: testdataStream10Lang1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Exec("DELETE FROM stream_speakers WHERE stream_id = ?", testdataStream10).Error; err != nil {
					return err
				}
				if err := tx.Exec("DELETE FROM stream_tags WHERE stream_id = ?", testdataStream10).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Stream{
					ID: testdataStream10,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "10-data-0030-01-stream-11",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataStream11,
					ChannelID: TestChannelID1,
					EventID:   &TestEventID2,
					Chat:      false,
					Running:   true,
					StartAt:   time.Date(2021, 7, 23, 17, 0, 0, 0, loc),
					ListenAt:  time.Date(2021, 7, 23, 17, 0, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/710b445a-51f5-4a9c-8fd4-59956453401c/poster.png",
					Tags: []*Tag{
						{ID: TestTagKonzertID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataStream11Speaker1,
							Name:    "Motschmann Trio",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataStream11Lang1,
					StreamID: testdataStream11,
					Lang:     "de",
					Title:    "Motschmann Trio",
					Subtitle: "Freitag ab 18:00",
					Short:    `Mit der Premiere der Electric Fields in der Berghain Kantine startete das Trio im Jahr 2016 seine Mission`,
					Long: `Mit der Premiere der Electric Fields in der Berghain Kantine startete das Trio im Jahr 2016 seine Mission: Elektronische Musik mit Originalinstrumenten so auf die Bühne zu bringen, dass alle Sounds live generiert werden. Mit einem riesigen Arsenal von analogen Synthesizern, E-Pianos und einem komplexen Multipercussion-Setup reisen Johannes Motschmann, Boris Bolles und David Panzl zwischen den Welten und treten in Clubs wie klassischen Konzertsälen gleichermaßen auf. Electric Fields ist ein Soundtrack, der von leisen Klavierklängen bis zu harten polymetrischen Beats reicht. Ambient- und Dronesounds wechseln sich mit orchestral anmutenden Passagen ab.

Mit einem alten Wurlitzerpiano, das Johannes und Boris im Keller ihres Studentenwohnheims aufgetan hatten, begann die Suche nach immer neuen Instrumenten, die den elektrischen Feldern Jahre später Klang und Gestalt gegeben haben. Im Zentrum stehen die Klänge des Wurlitzerpianos und der CP-70, die mit Bassklängen von Moog Prodigy und MS-20 das harmonische Fundament bilden. Ein gewisser retrospektiver Sound entsteht dadurch, dass fast alle Instrumente aus den 70er und 80er Jahren stammen.

David erweckt Rhythmen zum Leben, die vorab auf Drumcomputern konzipiert wurden und taucht sie in ein neues Licht, während Boris immer wieder mit der Violine einen zerbrechlichen Klang findet, der die Rhythmen und Pattern in einen sphärischen Sound führt. Alles was maschinengesteuert war, liegt nun wieder in den Händen der drei klassisch ausgebildeten Musiker, die mit hoher Präzision Johannes' akribisch ausnotierte Kompositionen so symphonisch klingen lassen, als würde man nicht einem elektroakustischen Trio sondern einem ganzen Orchester zuhören.`,
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Delete(&StreamLang{
					ID: testdataStream11Lang1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Exec("DELETE FROM stream_speakers WHERE stream_id = ?", testdataStream11).Error; err != nil {
					return err
				}
				if err := tx.Exec("DELETE FROM stream_tags WHERE stream_id = ?", testdataStream11).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Stream{
					ID: testdataStream11,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			ID: "10-data-0030-01-stream-12",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.Create(&Stream{
					ID:        testdataStream12,
					ChannelID: TestChannelID1,
					EventID:   &TestEventID2,
					Chat:      false,
					Running:   true,
					StartAt:   time.Date(2021, 7, 23, 20, 15, 0, 0, loc),
					ListenAt:  time.Date(2021, 7, 23, 20, 15, 0, 0, loc),
					Poster:    "https://cdn.media.kukoon.de/videos/df1555f5-7046-4f7a-adcc-195b73949723/54ff055b-5e46-4344-a43f-deb41c693045/poster.png",
					Tags: []*Tag{
						{ID: TestTagKonzertID},
					},
					Speakers: []*Speaker{
						{
							OwnerID: TestChannelID1,
							ID:      testdataStream12Speaker1,
							Name:    "Niklas Paschburg",
						},
					},
				}).Error; err != nil {
					return err
				}
				if err := tx.Create(&StreamLang{
					ID:       testdataStream12Lang1,
					StreamID: testdataStream12,
					Lang:     "de",
					Title:    "Niklas Paschburg",
					Subtitle: "Freitag ab 20:30",
					Short:    `2018 legte **Niklas Paschburg** sein Debütalbum Oceanic vor. Niklas‘ musikalische Ideen setzen sich aus einer Kombination umfangreicher Live-Performance sowie Experimenten zusammen. Im Februar 2020 erschien sein zweites Album Svalbard auf dem Label !7K.`,
					Long: `2018 legte **Niklas Paschburg** sein Debütalbum Oceanic vor. Niklas‘ musikalische Ideen setzen sich aus einer Kombination umfangreicher Live-Performance sowie Experimenten zusammen. Im Februar 2020 erschien sein zweites Album Svalbard auf dem Label !7K.

Niklas’ Gebrauch des Klaviers, elektronischer Elemente (Synth und Computer) sowie das Klavierakkordeon seines Großvaters ermöglichen es ihm, mit einer großen Bandbreite musikalischer Traditionen und Sprachen zu kommunizieren. Paschburgs Musik ist einzigartig in ihrer Fähigkeit, sowohl melancholisch als auch positiv zu sein — eine Umarmung, die Ängste abbaut und zur Meditation anregt, während sie die Hörer*in gleichzeitig dazu bringt, sich zur Musik zu bewegen, zu tanzen und zu rennen.

An dieser Stelle sei auch noch eine Vorstellung und ein Interview aus Pianoo.de die zutreffenste Beschreibung dieses großartigen Künstlers:
"**Niklas Paschburg** gehört zu den modernen Pianisten, die im weit gesteckten Bereich der Neo Klassik den Piano-Sound in den Mittelpunkt stellen. Das spannende Element, das Niklas’ Alben verbindet: Er beginnt ganz puristisch mit dem Klavierklang, um dann seine Zuhörer mitzunehmen auf eine klangmalerische Reise durch tiefe Soundscapes mit Clubtauglicher Intensität. Musik zum Wegfliegen und Wegtanzen.
Bei **Niklas Paschburg** funktioniert das nicht nur im Studio, sondern auch live. Mit Sequencing, Synthis und live gespielten Trommeln und Akkordeon schichtet er Motiv um Motiv. Dabei spielt er sehr gekonnt mit dem Kontrast zwischen fettem Synth-Orchester und sanften Sounds mit fragilem Anklang.

**Klavier- oder Elektronische Musik?**
Beides trifft zu. Ein perfektes Beispiel dafür ist das 2019 veröffentlichte Stück »Blooming in C minor«. Niklas greift hier das Präludium von J.S. Bach auf. Wer jetzt gleich altmodische Synthesizer-Musik á la »Switched On Bach« von Wendy Carlos vermutet – weit gefehlt. Niklas geht mit dem Material ganz anders um. Auf seine Weise eben. Er zerlegt das Präludium komplett, wiederholt bestimmte Passagen und lässt sie ineinander laufen. Die so entstehenden changierenden Farben stehen schon in einem Widerspruch zu der klaren Kompositionslinie Bachs.
Niklas: Das Stück habe ich – ja klar – früher in seiner klassischen Form gespielt und geübt. Dann hatte ich es wiederentdeckt, wobei dann aber der Gedanke entstand, vor allem live etwas damit zu machen. Es ist ein starker Live-Song daraus geworden.
Ist das dann ein festgelegter Song, den du immer identisch live spielst oder arbeitest du mit improvisatorischen Elementen?
Niklas: Es ist immer beides. Ich habe eine Struktur, die schon feststeht. Aber ich kann die einzelnen Elemente immer variieren. Ich kann die Teile verlängern, ich kann eine neue Melodie drüber spielen, ein Intro oder Outro ganz anders machen. Also die Sachen, die ich einloope, sind erstmal geplant. Klar, kann da immer mal etwas Unvorhergesehenes passieren und ich gehe dann in eine ganz andere Richtung, aber im Prinzip stehen diese Elemente fest.
 
**Live-Arrangement vs. Improvisation**
Deine Musik hat immer etwas sehr Spontanes – und dennoch wirkt es arrangiert. Und es ist beeindruckend, welchen Sound du ganz alleine auf der Bühne entwickelst. Vor allem auch in der Spieldynamik, wobei auch abrupte Wechsel vom fetten Arrangement hin zu einem ganz minimalistischen Part möglich sind.
Niklas: Es gibt grundsätzlich die Möglichkeit, alles zu improvisieren, wo du dann bei Null oder mit einer Grundidee beginnst. Mir war aber von vornherein wichtig, in Song-Strukturen zu arbeiten. Da ich so viele Instrumente an bestimmten Punkten reinkommen lassen will, war gleich klar, dass ich das schon durchplanen muss. Oder andersherum beschrieben: Ohne eine Planung des Ganzen müsste ich die Auswahl der Instrumente sehr einschränken, und das wiederum würde sich auf die Musik auswirken.
Sicher sehr stark vor allem auch auf die Dramaturgie der Musik.
Niklas: Ich will mit verschiedenen Parts arbeiten, also von einem C-Part wieder zurück auf den großen Part mit dem Akkordeon, der Bassdrum, Pads und allem, was dazu gehört. Ohne eine Planung wäre man da wohl überfordert.
 
**Niklas Paschburg live**
Wie bringst du das alles auf die Bühne?
Niklas: Das Herz vom ganzen Live-Setup ist Ableton Live. Dort habe ich die Strukturen der Songs als Patterns und Scenes angelegt. Das ermöglicht es mir, die einzelnen Parts der Songs gezielt anzusteuern. Ich kann spontan einen Part überspringen, kann jederzeit auf den Punkt genau zu einem bestimmten Part zurückkommen. Das kann ich auf der Bühne frei entscheiden.
Nutzt du Ableton Live auch für das Live-Looping, um z.B. die Bassdrum aufzunehmen und als Loops im Arrangement zu verankern.
Niklas: Ja, genau. Es gibt zwar auch gute externe Looper, und dann brauchst du gar keinen Computer oder Laptop. Mir war es aber wichtig, dass ich jedes Signal einzeln ausspielen kann. Und das ist bei den meisten dieser Geräte nicht möglich. Außerdem ist es mir wesentlich lieber, alles in einem zentralen Projekt zu haben.
Beim Live-Looping muss man immer auch der Versuchung widerstehen, zu viele Elemente zu schichten. Wie hältst du die Balance dabei?
Niklas: (lacht) Ich kämpfe dabei immer so ein bisschen. Man hat immer noch eine und noch eine Idee und weiß am Ende nicht, wo man es unterbringen soll. Aber für mich hat das viel mit Ausprobieren zu tun. Damit beginne ich in meinem Studio. Ich versuche dann immer abzuwägen: Sind es vielleicht zu viele Parts, sodass ich mich schon fast um nichts anderes mehr kümmern kann beim Spielen? Habe ich dann noch die Möglichkeit, das Klavier zu spielen? Irgendwann nehme ich das Stück dann mit auf die Bühne, um es weiter zu testen. Stelle ich dann fest, dass alles zu viel ist, verwerfe ich auch Dinge und suche nach einer neuen Lösung.
 
**Live-Location und Dramaturgie**
Aber man stelle sich vor – alles ist perfekt vorbereitet, du kommst auf die Bühne und stellst spontan fest, dass die Intensität der Musik nicht zur Stimmung in der Location passt. Kommt so etwas vor und kannst du dann darauf eingehen?
Niklas: Tatsächlich kommt so etwas vor, und ich bin darauf vorbereitet. Ich spiele in sehr unterschiedlichen Locations – einmal in Kirchen, wo es sehr atmosphärisch und ruhiger ist. Dann wiederum spiele ich in Clubs, wo dann die Leute stehen und auch tanzen können. Das hat viel mehr Energie …
… und erfordert eine andere Intensität.
Niklas: Genau. So existieren im Grunde zwei Varianten des Live-Sets. Aber ich spiele auch anders. Ist die Atmosphäre ruhiger, dann spiele ich auch die Bassdrum sanfter, spiele sanfter Klavier und reize auch die Synthesizer nicht bis an ihre Grenzen aus. Im Club weiß ich dann, dass ich Gas geben und richtig laut sein kann. Auch die Auswahl und Reihenfolge der Songs kann sich spontan ändern, wenn ich das Gefühl habe, dass etwas nicht passt.
Du nutzt sehr unterschiedliches Equipment – wo würdest du dich selber einordnen: Elektronik oder Piano?
Niklas: Das ist sehr schwierig. Und im Grunde möchte ich mich da gar nicht festlegen. Denn es ändert sich ständig – mal ist für mich das Klavier das Instrument. Es ist dann absolut im Mittelpunkt und ich bastle alle andere Sounds drumherum. Bei manchen Songs aber spielt das Klavier dann gar nicht mehr die Hauptrolle. Nimm z.B. »Blooming in C minor« – das würde ich als Elektronische Musik bezeichnen.
Aber – wie bei vielen anderen deiner Tracks auch – mit vielen akustischen Elementen. Das Klavier nutzt du ja für rhythmische Patterns, Texturen und Sequenz-Motive und oben drauf dann noch mal eine Melodie usw.
Niklas: Das stimmt schon – mir sind die akustischen Sounds einfach sehr wichtig. Ich wollte auch unbedingt eine akustische Bassdrum haben, da ich das Gefühl habe, dass das lebt, da ich das Instrument spontan so unterschiedlich spielen kann.
 
**Was ist für dich das Besondere am Klavier?**
Niklas: Das Klavier ist ein vielseitiges Instrument mit einem wahnsinnig breiten Spektrum. Du kannst so viele verschiedene Sachen damit machen – du kannst es zupfen, leise und laut spielen, hart und sanft spielen. Und dann hat es diesen Full-Range-Tonumfang von 88 Tasten. Ich entdecke immer wieder etwas Neues.
Bei den Synthesizern gibt es dann ja die Frage: Hardware oder Software?
Niklas: Live verwende ich hauptsächlich die Hardware-Synthis: einen OB-6 von Dave Smith Instruments und einen Arp Odyssey als Lead-Synth. Als Software-Instrument nutze ich eigentlich nur noch Kontakt von Native Instruments. Dafür habe ich mir mal ein eigenes Instrument aus Samples gebastelt. Dann sind da noch ein paar einzelne Samples, die ich mit dem OB6 und einem Korg MS2000 aufgenommen habe. Es ist ja bei den Synthis immer auch die Diskussion, ob nun Hardware oder Software besser klingt. Das ist für mich aber gar nicht der Punkt. Es ist einfach so, dass ich beim Entwickeln von Sounds etwas Handfestes zum Anfassen haben muss. Also: ja, bei Synthesizern bin ich schon eher der Hardware-Typ.

**Wenn du live spielst – strebst du dann einen bestimmten Sound an oder einen Zustand?**
Niklas: Ganz klar: Beim Live-Spielen geht es mir darum, einen Zustand zu erreichen. Jeder soll die Möglichkeit haben, auf seine eigene Reise zu gehen. Klar spielt der Sound immer eine große Rolle, und er ermöglicht es mir auch auf gewisse Weise, einen Zustand zu kreieren. Und dabei wiederum kommt es sehr darauf an, wo ich gerade spiele – in einer Kirche oder in einem Club. In einer Kirche möchte ich dann eher einen Zustand kreieren, wo die Leute in eine Atmosphäre eintauchen und ihren Gedanken freien Lauf lassen können. Im Club kann ich dann eine Richtung vorgeben, wo es etwas tanzbarer werden darf.
Niklas, vielen Dank für das Interview"
[Quelle: [www.pianoo.de/pianoo-people/niklas-paschburg](https://www.pianoo.de/pianoo-people/niklas-paschburg/)]
`,
				}).Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Delete(&StreamLang{
					ID: testdataStream12Lang1,
				}).Error; err != nil {
					return err
				}
				if err := tx.Exec("DELETE FROM stream_speakers WHERE stream_id = ?", testdataStream12).Error; err != nil {
					return err
				}
				if err := tx.Exec("DELETE FROM stream_tags WHERE stream_id = ?", testdataStream12).Error; err != nil {
					return err
				}
				if err := tx.Delete(&Stream{
					ID: testdataStream12,
				}).Error; err != nil {
					return err
				}
				return nil
			},
		},
	}...)
}
