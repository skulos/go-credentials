package commands

import (
	"fmt"
	"log"
	"os"
	"sort"

	"filippo.io/age"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/skulos/go-credentials/internal/crypto"
	"github.com/skulos/go-credentials/internal/environment"
	"gopkg.in/yaml.v3"
)

// SaveUpdatedCredentials takes the updated credentials map and saves it to a YAML file
func SaveUpdatedCredentials(env, content string, id *age.X25519Identity) error {

	encName := environment.ResolveEnv(env, false)
	encPath := fmt.Sprintf(".credentials/%s.yml.enc", encName)

	// Parse the edited YAML content to ensure it's valid
	var data map[string]interface{}
	if err := yaml.Unmarshal([]byte(content), &data); err != nil {
		return fmt.Errorf("failed to parse edited credentials: %w", err)
	}

	// Marshal it again to ensure clean formatting
	newYAML, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal updated credentials: %w", err)
	}

	// Encrypt and overwrite the encrypted file
	recipient := id.Recipient()
	if err := crypto.EncryptToFile(encPath, recipient, newYAML); err != nil {
		return fmt.Errorf("failed to encrypt and save updated credentials: %w", err)
	}

	return nil
}

func Editor(env string) (string, bool, error) {

	var edited bool = false

	keyName := environment.ResolveEnv(env, true)
	encName := environment.ResolveEnv(env, false)

	keyPath := fmt.Sprintf(".credentials/%s.key", keyName)
	encPath := fmt.Sprintf(".credentials/%s.yml.enc", encName)

	keyBytes, err := os.ReadFile(keyPath)
	if err != nil {
		return "", edited, fmt.Errorf("failed to read private key: %w", err)
	}

	identity, err := crypto.ParseIdentity(string(keyBytes))
	if err != nil {
		return "", edited, fmt.Errorf("failed to parse identity: %w", err)
	}

	plaintext, err := crypto.DecryptFromFile(encPath, identity)
	if err != nil {
		return "", edited, fmt.Errorf("failed to decrypt credentials: %w", err)
	}

	var data map[string]interface{}
	if err := yaml.Unmarshal(plaintext, &data); err != nil {
		return "", edited, fmt.Errorf("failed to parse YAML: %w", err)
	}

	// Pretty print top-level keys (flat output)
	keys := make([]string, 0, len(data))
	maxKeyLen := 0
	for k := range data {
		keys = append(keys, k)
		if len(k) > maxKeyLen {
			maxKeyLen = len(k)
		}
	}

	sort.Strings(keys)

	yamlBytes, err := yaml.Marshal(data)
	if err != nil {
		return "", edited, fmt.Errorf("failed to format YAML: %w", err)
	}

	app := tview.NewApplication()

	// Create a plain TextArea (no syntax highlighting)
	textArea := tview.NewTextArea()
	textArea.SetText(string(yamlBytes), false)
	textArea.SetTitle(" Editing " + encPath + " ").SetBorder(true)

	position := tview.NewTextView().SetDynamicColors(true).SetTextAlign(tview.AlignRight)
	updateInfos := func() {
		fromRow, fromColumn, toRow, toColumn := textArea.GetCursor()
		if fromRow == toRow && fromColumn == toColumn {
			position.SetText(fmt.Sprintf("Row: [yellow]%d[white], Column: [yellow]%d", fromRow, fromColumn))
		} else {
			position.SetText(fmt.Sprintf("[red]From[white] Row: [yellow]%d[white], Column: [yellow]%d[white] - [red]To[white] Row: [yellow]%d[white], To Column: [yellow]%d", fromRow, fromColumn, toRow, toColumn))
		}
	}
	textArea.SetMovedFunc(updateInfos)
	updateInfos()

	// Help text for Save and Cancel
	helpInfo := tview.NewTextView().
		SetText("Press [yellow]Ctrl + S[white] to [green]save, [yellow]Ctrl + C[white] to [red]cancel/exit").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	// Create a main view with a grid layout
	mainView := tview.NewGrid().
		SetRows(0, 1).
		AddItem(textArea, 0, 0, 1, 2, 0, 0, true).
		AddItem(position, 1, 1, 1, 1, 0, 0, false).
		AddItem(helpInfo, 1, 0, 1, 2, 0, 0, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlS:
			editedContent := textArea.GetText()
			err := SaveUpdatedCredentials(env, editedContent, identity)
			if err != nil {
				log.Fatalf("Failed to save updated credentials: %v", err)
			}
			app.Stop()
			edited = true
			return nil
		case tcell.KeyCtrlC:
			app.Stop()
			return nil
		}
		return event
	})

	// Run the application
	if err := app.SetRoot(mainView, true).EnableMouse(true).Run(); err != nil {
		log.Fatalf("Error starting application: %v", err)
	}

	return encPath, edited, nil
}
