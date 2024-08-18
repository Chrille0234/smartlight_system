package groups

import (
    "encoding/json"
    "fmt"
    "os"
    "path/filepath"
    "hue/types"
)

func Add_group(group_name string, ids map[string]string) error {
    // Define the path to the groups.json file
    filePath := filepath.Join("database", "groups", "groups.json")

    // Read the existing groups from the file
    file, err := os.Open(filePath)
    if err != nil {
        return fmt.Errorf("failed to open groups.json: %v", err)
    }
    defer file.Close()

    var groups types.Groups
    err = json.NewDecoder(file).Decode(&groups)
    if err != nil {
        return fmt.Errorf("failed to decode groups.json: %v", err)
    }

    // Add the new group
    groups.Groups[group_name] = types.Group{IDs: ids}

    // Write the updated groups back to the file
    fileData, err := json.MarshalIndent(groups, "", "  ")
    if err != nil {
        return fmt.Errorf("failed to marshal groups: %v", err)
    }

    err = os.WriteFile(filePath, fileData, 0644)
    if err != nil {
        return fmt.Errorf("failed to write to groups.json: %v", err)
    }

    return nil
}
