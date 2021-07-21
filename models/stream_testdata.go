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
					Poster:    "https://kukoon.de/images/8/6/7/e/5/767e5a264900c670ff18777ee9d5e67466c2a185-martin-kohlstedt.jpg",
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
	}...)
}
