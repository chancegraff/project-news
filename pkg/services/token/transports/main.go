package transports

import (
	"time"

	pb "github.com/chancegraff/project-news/api/token"
	"github.com/chancegraff/project-news/internal/models"
)

// DecodeProtoClient will decode a protobuff Client into the database model
func DecodeProtoClient(pbc *pb.Client) models.Client {
	layout := "2006-01-02 15:04:05 -0700 MST"
	clientExpiration, _ := time.Parse(layout, pbc.ExpiredAt)
	return models.Client{
		Hash:      pbc.Hash,
		UserID:    pbc.UserID,
		IP:        pbc.Ip,
		ExpiredAt: clientExpiration,
	}
}

// DecodeProtoIdentifiers will decode a protobuff Identifier into the database model
func DecodeProtoIdentifiers(pbi *pb.Identifiers) models.Identifiers {
	return models.Identifiers{
		Software: pbi.Software,
		Browser:  pbi.Browser,
		Language: pbi.Language,
		Width:    pbi.Width,
		Height:   pbi.Height,
		Colors:   pbi.Colors,
		Pixels:   pbi.Pixels,
	}
}
