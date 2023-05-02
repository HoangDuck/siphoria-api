package controller

import (
	"github.com/labstack/echo/v4"
	response "hotel-booking-api/model/model_func"
)

type MapModel struct {
	Value string `json:"data"`
}

func GetEmbeddedMap(c echo.Context) error {
	value := c.Param("id")
	tempMapInfo := ""
	switch value {
	case "bv":
		tempMapInfo = `<iframe src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d2013477.3230151867!2d105.73726910266026!3d9.717445107324528!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x31756c9c282e8e43%3A0xcce3539941eaed52!2zQsOgIFLhu4thIC0gVsWpbmcgVMOgdSwgVmnhu4d0IE5hbQ!5e0!3m2!1svi!2s!4v1683036227167!5m2!1svi!2s" width="600" height="450" style="border:0;" allowfullscreen="" loading="lazy" referrerpolicy="no-referrer-when-downgrade"></iframe>`
	case "bg":
		tempMapInfo = `<iframe src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d59479.79792778326!2d106.14884955624593!3d21.291750191570653!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x31356dadb70fbfe5%3A0xd6dbe565b8b15e5c!2zVHAuIELhuq9jIEdpYW5nLCBC4bqvYyBHaWFuZywgVmnhu4d0IE5hbQ!5e0!3m2!1svi!2s!4v1683036272101!5m2!1svi!2s" width="600" height="450" style="border:0;" allowfullscreen="" loading="lazy" referrerpolicy="no-referrer-when-downgrade"></iframe>`
	case "bk":
		tempMapInfo = `<iframe src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d59479.79792778326!2d106.14884955624593!3d21.291750191570653!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x31356dadb70fbfe5%3A0xd6dbe565b8b15e5c!2zVHAuIELhuq9jIEdpYW5nLCBC4bqvYyBHaWFuZywgVmnhu4d0IE5hbQ!5e0!3m2!1svi!2s!4v1683036272101!5m2!1svi!2s" width="600" height="450" style="border:0;" allowfullscreen="" loading="lazy" referrerpolicy="no-referrer-when-downgrade"></iframe>`
	case "bl":
		tempMapInfo = `<iframe src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d126007.30962094327!2d105.67055675232942!3d9.26854252051459!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x31a10a2351f087b3%3A0x4949992f9e65b750!2zVHAuIELhuqFjIExpw6p1LCBC4bqhYyBMacOqdSwgVmnhu4d0IE5hbQ!5e0!3m2!1svi!2s!4v1683036365196!5m2!1svi!2s" width="600" height="450" style="border:0;" allowfullscreen="" loading="lazy" referrerpolicy="no-referrer-when-downgrade"></iframe>`
	case "bn":
		tempMapInfo = `<iframe src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d119054.53574404765!2d105.92191312582536!3d21.17410682235044!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x31350c5b3464ae51%3A0x1a3035b9749102f9!2zVHAuIELhuq9jIE5pbmgsIELhuq9jIE5pbmgsIFZp4buHdCBOYW0!5e0!3m2!1svi!2s!4v1683036430770!5m2!1svi!2s" width="600" height="450" style="border:0;" allowfullscreen="" loading="lazy" referrerpolicy="no-referrer-when-downgrade"></iframe>`
	case "bt":
		tempMapInfo = `<iframe src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d62820.77908641769!2d106.33426581282512!3d10.237476314793662!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x310aa8f5e2e8bd09%3A0x9d5fd18ce4fa56bb!2zVHAuIELhur9uIFRyZSwgQuG6v24gVHJlLCBWaeG7h3QgTmFt!5e0!3m2!1svi!2s!4v1683036479519!5m2!1svi!2s" width="600" height="450" style="border:0;" allowfullscreen="" loading="lazy" referrerpolicy="no-referrer-when-downgrade"></iframe>`
	case "bd":
		tempMapInfo = `<iframe src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d990604.6774731488!2d108.31883665129868!3d14.103749374446107!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x316f2def6e711bbf%3A0x45bf4c043ae5fd37!2zQsOsbmggxJDhu4tuaCwgVmnhu4d0IE5hbQ!5e0!3m2!1svi!2s!4v1683036510518!5m2!1svi!2s" width="600" height="450" style="border:0;" allowfullscreen="" loading="lazy" referrerpolicy="no-referrer-when-downgrade"></iframe>`
	default:
		tempMapInfo = `<iframe src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d990604.6774731488!2d108.31883665129868!3d14.103749374446107!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x316f2def6e711bbf%3A0x45bf4c043ae5fd37!2zQsOsbmggxJDhu4tuaCwgVmnhu4d0IE5hbQ!5e0!3m2!1svi!2s!4v1683036510518!5m2!1svi!2s" width="600" height="450" style="border:0;" allowfullscreen="" loading="lazy" referrerpolicy="no-referrer-when-downgrade"></iframe>`
	}
	return response.Ok(c, "success", MapModel{
		Value: tempMapInfo,
	})
}
