package handlers

import (
	"context"

	ilsp "github.com/hashicorp/terraform-ls/internal/lsp"
	lsp "github.com/hashicorp/terraform-ls/internal/protocol"
)

func (svc *service) SignatureHelp(ctx context.Context, params lsp.SignatureHelpParams) (*lsp.SignatureHelp, error) {
	_, err := ilsp.ClientCapabilities(ctx)
	if err != nil {
		return nil, err
	}

	dh := ilsp.HandleFromDocumentURI(params.TextDocument.URI)
	doc, err := svc.stateStore.DocumentStore.GetDocument(dh)
	if err != nil {
		return nil, err
	}

	d, err := svc.decoderForDocument(ctx, doc)
	if err != nil {
		return nil, err
	}

	pos, err := ilsp.HCLPositionFromLspPosition(params.Position, doc)
	if err != nil {
		return nil, err
	}

	sig, err := d.SignatureAtPos(doc.Filename, pos)
	if err != nil {
		return nil, err
	}

	return ilsp.ToSignatureHelp(sig), nil
}
