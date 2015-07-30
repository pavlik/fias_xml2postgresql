package address_object

import (
	"encoding/xml"

	"github.com/pavlik/fias_xml2postgresql/helpers"
)

func InitAdressObject() helpers.ExportParams {
	tableName := "as_addrobj"
	elementName := "Object"

	// Классификатор адресообразующих элементов
	type xmlObject struct {
		XMLName    xml.Name `xml:"Object"`
		AOGUID     string   `xml:"AOGUID,attr"`
		FORMALNAME string   `xml:"FORMALNAME,attr"`
		REGIONCODE int      `xml:"REGIONCODE,attr"`
		AUTOCODE   int      `xml:"AUTOCODE,attr"`
		AREACODE   int      `xml:"AREACODE,attr"`
		CITYCODE   int      `xml:"CITYCODE,attr"`
		CTARCODE   int      `xml:"CTARCODE,attr"`
		PLACECODE  int      `xml:"PLACECODE,attr"`
		STREETCODE int      `xml:"STREETCODE,attr,omitempty"`
		EXTRCODE   int      `xml:"EXTRCODE,attr"`
		SEXTCODE   int      `xml:"SEXTCODE,attr"`
		OFFNAME    *string  `xml:"OFFNAME,attr,omitempty"`
		POSTALCODE *string  `xml:"POSTALCODE,attr,omitempty"`
		IFNSFL     int      `xml:"IFNSFL,attr,omitempty"`
		TERRIFNSFL int      `xml:"TERRIFNSFL,attr,omitempty"`
		IFNSUL     int      `xml:"IFNSUL,attr,omitempty"`
		TERRIFNSUL int      `xml:"TERRIFNSUL,attr,omitempty"`
		OKATO      *string  `xml:"OKATO,attr,omitempty"`
		OKTMO      *string  `xml:"OKTMO,attr,omitempty"`
		UPDATEDATE string   `xml:"UPDATEDATE,attr"`
		SHORTNAME  string   `xml:"SHORTNAME,attr"`
		AOLEVEL    int      `xml:"AOLEVEL,attr"`
		PARENTGUID *string  `xml:"PARENTGUID,attr,omitempty"`
		AOID       string   `xml:"AOID,attr"`
		PREVID     *string  `xml:"PREVID,attr,omitempty"`
		NEXTID     *string  `xml:"NEXTID,attr,omitempty"`
		CODE       *string  `xml:"CODE,attr,omitempty"`
		PLAINCODE  *string  `xml:"PLAINCODE,attr,omitempty"`
		ACTSTATUS  bool     `xml:"ACTSTATUS,attr"`
		CENTSTATUS bool     `xml:"CENTSTATUS,attr"`
		OPERSTATUS int      `xml:"OPERSTATUS,attr"`
		CURRSTATUS int      `xml:"CURRSTATUS,attr"`
		STARTDATE  string   `xml:"STARTDATE,attr"`
		ENDDATE    string   `xml:"ENDDATE,attr"`
		NORMDOC    *string  `xml:"NORMDOC,attr,omitempty"`
		LIVESTATUS bool     `xml:"LIVESTATUS,attr"`
	}

	schema := `CREATE TABLE ` + tableName + ` (
	    ao_guid UUID NOT NULL,
	    formal_name VARCHAR(120) NOT NULL,
			region_code INT NOT NULL,
			auto_code INT NOT NULL,
			area_code INT NOT NULL,
			city_code INT NOT NULL,
			ctar_code INT NOT NULL,
			place_code INT NOT NULL,
			street_code INT,
			extr_code INT NOT NULL,
			sext_code INT NOT NULL,
			off_name VARCHAR(120),
			postal_code VARCHAR(6),
			ifns_fl INT,
			terr_ifns_fl INT,
			ifns_ul INT,
			terr_ifns_ul INT,
			okato VARCHAR(11),
			oktmo VARCHAR(11),
			update_date TIMESTAMP NOT NULL,
			short_name VARCHAR(10) NOT NULL,
			ao_level INT NOT NULL,
			parent_guid UUID,
			ao_id UUID NOT NULL,
			prev_id UUID,
			next_id UUID,
			code VARCHAR(17),
			plain_code VARCHAR(15),
			act_status BOOL NOT NULL,
			cent_status BOOL NOT NULL,
			oper_status INT NOT NULL,
			curr_status INT NOT NULL,
			start_date TIMESTAMP NOT NULL,
			end_date TIMESTAMP NOT NULL,
			norm_doc UUID,
			live_status BOOL NOT NULL,
			PRIMARY KEY (ao_id));`

	ao := new(helpers.ExportParams)

	ao.tableName = tableName
	ao.elementName = elementName
	ao.schema = schema
	ao.xmlObject = xmlObject
	ao.fields = []string{"ao_guid",
		"formal_name",
		"region_code",
		"auto_code",
		"area_code",
		"city_code",
		"ctar_code",
		"place_code",
		"street_code",
		"extr_code",
		"sext_code",
		"off_name",
		"postal_code",
		"ifns_fl",
		"terr_ifns_fl",
		"ifns_ul",
		"terr_ifns_ul",
		"okato",
		"oktmo",
		"update_date",
		"short_name",
		"ao_level",
		"parent_guid",
		"ao_id",
		"prev_id",
		"next_id",
		"code",
		"plain_code",
		"act_status",
		"cent_status",
		"oper_status",
		"curr_status",
		"start_date",
		"end_date",
		"norm_doc",
		"live_status"}

	return ao
}