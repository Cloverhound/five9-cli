#!/usr/bin/env python3
"""Generate cmd/*.go files from five9_api_spec.json."""

import json
import os
import re
import sys

SPEC_PATH = os.path.join(os.path.dirname(__file__), "five9_api_spec.json")
CMD_DIR = os.path.join(os.path.dirname(__file__), "..", "cmd")


def to_go_var(s):
    """Convert kebab-case CLI command name to Go variable name."""
    parts = s.split("-")
    return "".join(p.capitalize() for p in parts)


def to_go_filename(s):
    """Convert CLI command name to Go filename."""
    return s.replace("-", "_") + ".go"


def method_to_subcommand(method_name, cli_cmd):
    """Convert a SOAP method name to a CLI subcommand name.

    Examples:
        getAgentGroups -> list (for agent-groups command)
        createAgentGroup -> create
        deleteAgentGroup -> delete
        modifyAgentGroup -> modify
        getAgentGroup -> get
        addDispositionsToCampaign -> add-dispositions
        userSkillAdd -> skill-add
    """
    # Build a map of known patterns
    name = method_name

    # Special cases for specific patterns
    special = {
        # Reports
        "runReport": "run",
        "isReportRunning": "status",
        "getReportResult": "get",
        "getReportResultCsv": "get-csv",
        # VCC
        "getApiVersions": "get-api-versions",
        "getVCCConfiguration": "get-config",
        "modifyVCCConfiguration": "modify-config",
        # Session
        "closeSession": "close",
        # Domain limits
        "getCallCountersState": "get-call-counters",
        # Lists
        "getListsInfo": "list-info",
        "getListImportResult": "import-status",
        "isImportRunning": "is-importing",
        "addNumbersToDnc": "add-numbers-to-dnc",
        "removeNumbersFromDnc": "remove-numbers-from-dnc",
        "checkDncForNumbers": "check-dnc",
        "addRecordToList": "add-record",
        "addRecordToListSimple": "add-record-simple",
        "addToList": "add-records",
        "addToListCsv": "add-records-csv",
        "addToListFtp": "add-records-ftp",
        "asyncAddRecordsToList": "async-add-records",
        "deleteFromList": "delete-records",
        "deleteFromListCsv": "delete-records-csv",
        "deleteFromListFtp": "delete-records-ftp",
        "asyncDeleteRecordsFromList": "async-delete-records",
        "deleteRecordFromList": "delete-record",
        "deleteAllFromList": "delete-all",
        "createList": "create",
        "deleteList": "delete",
        # Contacts
        "getContactRecords": "get-records",
        "updateCrmRecord": "update-crm-record",
        "updateContacts": "update",
        "updateContactsCsv": "update-csv",
        "updateContactsFtp": "update-ftp",
        "asyncUpdateCrmRecords": "async-update",
        "deleteFromContacts": "delete",
        "deleteFromContactsCsv": "delete-csv",
        "deleteFromContactsFtp": "delete-ftp",
        "getCrmImportResult": "import-status",
        "getContactFields": "get-fields",
        "createContactField": "create-field",
        "modifyContactField": "modify-field",
        "deleteContactField": "delete-field",
        # Campaigns (configuration)
        "getCampaigns": "list",
        "createInboundCampaign": "create-inbound",
        "createOutboundCampaign": "create-outbound",
        "createAutodialCampaign": "create-autodial",
        "deleteCampaign": "delete",
        "getInboundCampaign": "get-inbound",
        "getOutboundCampaign": "get-outbound",
        "getAutodialCampaign": "get-autodial",
        "modifyInboundCampaign": "modify-inbound",
        "modifyOutboundCampaign": "modify-outbound",
        "modifyAutodialCampaign": "modify-autodial",
        "renameCampaign": "rename",
        "addDispositionsToCampaign": "add-dispositions",
        "removeDispositionsFromCampaign": "remove-dispositions",
        "addDNISToCampaign": "add-dnis",
        "removeDNISFromCampaign": "remove-dnis",
        "addListsToCampaign": "add-lists",
        "removeListsFromCampaign": "remove-lists",
        "addSkillsToCampaign": "add-skills",
        "removeSkillsFromCampaign": "remove-skills",
        "getCampaignDNISList": "get-dnis-list",
        "getDNISList": "get-all-dnis",
        "getListsForCampaign": "get-lists",
        "setDefaultIVRSchedule": "set-default-ivr-schedule",
        "modifyCampaignLists": "modify-lists",
        # Campaign management
        "startCampaign": "start",
        "stopCampaign": "stop",
        "forceStopCampaign": "force-stop",
        "resetCampaign": "reset",
        "resetCampaignDispositions": "reset-dispositions",
        "resetListPosition": "reset-list-position",
        "getCampaignState": "get-state",
        "getCampaignStrategies": "get-strategies",
        "setCampaignStrategies": "set-strategies",
        "asyncUpdateCampaignDispositions": "async-update-dispositions",
        "updateDispositions": "update-dispositions",
        "updateDispositionsCsv": "update-dispositions-csv",
        "updateDispositionsFtp": "update-dispositions-ftp",
        "getDispositionsImportResult": "get-dispositions-import-result",
        # Campaign profiles
        "createCampaignProfile": "create",
        "deleteCampaignProfile": "delete",
        "getCampaignProfiles": "list",
        "getCampaignProfileDispositions": "get-dispositions",
        "getCampaignProfileFilter": "get-filter",
        "modifyCampaignProfile": "modify",
        "modifyCampaignProfileCrmCriteria": "modify-crm-criteria",
        "modifyCampaignProfileDispositions": "modify-dispositions",
        "modifyCampaignProfileFilterOrder": "modify-filter-order",
        # Skills
        "getSkills": "list",
        "getSkillsInfo": "list-info",
        "createSkill": "create",
        "deleteSkill": "delete",
        "getSkill": "get",
        "getSkillInfo": "get-info",
        "modifySkill": "modify",
        "addSkillAudioFile": "add-audio-file",
        "removeSkillAudioFile": "remove-audio-file",
        "getSkillAudioFiles": "get-audio-files",
        "getSkillVoicemailGreeting": "get-voicemail-greeting",
        "setSkillVoicemailGreeting": "set-voicemail-greeting",
        # Users
        "getUserInfo": "get",
        "getUsersInfo": "list-info",
        "getUsersGeneralInfo": "list",
        "getUserGeneralInfo": "get-general-info",
        "createUser": "create",
        "deleteUser": "delete",
        "modifyUser": "modify",
        "userSkillAdd": "skill-add",
        "userSkillModify": "skill-modify",
        "userSkillRemove": "skill-remove",
        "getUserVoicemailGreeting": "get-voicemail-greeting",
        "setUserVoicemailGreeting": "set-voicemail-greeting",
        "modifyUserCannedReports": "modify-canned-reports",
        # User profiles
        "getUserProfiles": "list",
        "getUserProfile": "get",
        "createUserProfile": "create",
        "deleteUserProfile": "delete",
        "modifyUserProfile": "modify",
        "modifyUserProfileSkills": "modify-skills",
        "modifyUserProfileUserList": "modify-user-list",
        # Agent groups
        "getAgentGroups": "list",
        "getAgentGroup": "get",
        "createAgentGroup": "create",
        "deleteAgentGroup": "delete",
        "modifyAgentGroup": "modify",
        # Dispositions
        "getDispositions": "list",
        "getDisposition": "get",
        "createDisposition": "create",
        "modifyDisposition": "modify",
        "removeDisposition": "remove",
        "renameDisposition": "rename",
        # Reason codes
        "createReasonCode": "create",
        "deleteReasonCode": "delete",
        "deleteReasonCodeByType": "delete-by-type",
        "getReasonCode": "get",
        "getReasonCodeByType": "get-by-type",
        "modifyReasonCode": "modify",
        # Web connectors
        "createWebConnector": "create",
        "deleteWebConnector": "delete",
        "getWebConnectors": "list",
        "modifyWebConnector": "modify",
        # Call variables
        "createCallVariable": "create",
        "deleteCallVariable": "delete",
        "getCallVariables": "list",
        "modifyCallVariable": "modify",
        "createCallVariablesGroup": "create-group",
        "deleteCallVariablesGroup": "delete-group",
        "getCallVariableGroups": "list-groups",
        "modifyCallVariablesGroup": "modify-group",
        # Speed dial
        "createSpeedDialNumber": "create",
        "getSpeedDialNumbers": "list",
        "removeSpeedDialNumber": "remove",
        # IVR scripts
        "createIVRScript": "create",
        "deleteIVRScript": "delete",
        "getIVRScripts": "list",
        "modifyIVRScript": "modify",
        "getIvrIcons": "get-icons",
        "setIvrIcons": "set-icons",
        "removeIvrIcons": "remove-icons",
        "getIvrScriptOwnership": "get-ownership",
        "setIvrScriptOwnership": "set-ownership",
        "removeIvrScriptOwnership": "remove-ownership",
        # Prompts
        "createPrompt": "create",
        "deletePrompt": "delete",
        "getPrompt": "get",
        "getPrompts": "list",
        "modifyPrompt": "modify",
        "addPromptWav": "add-wav",
        "addPromptWavInline": "add-wav-inline",
        "modifyPromptWav": "modify-wav",
        "modifyPromptWavInline": "modify-wav-inline",
        "addPromptTTS": "add-tts",
        "modifyPromptTTS": "modify-tts",
        "deleteLanguagePrompt": "delete-language",
        # Dialing rules
        "getDialingRules": "get",
        "setDialingRules": "set",
        # Locales
        "getAvailableLocales": "list-available",
        "getLocale": "get",
        "setLocale": "set",
        "getConfigurationTranslations": "get-translations",
        "updateConfigurationTranslations": "update-translations",
    }

    if name in special:
        return special[name]

    # Generic fallback: camelCase to kebab-case
    result = re.sub(r"([a-z])([A-Z])", r"\1-\2", name).lower()
    return result


def build_sample_json(type_name, types_dict, depth=0, seen=None):
    """Recursively build a sample JSON dict for a given type."""
    if seen is None:
        seen = set()
    if depth > 3 or type_name in seen:
        return {}
    seen = seen | {type_name}

    type_def = types_dict.get(type_name)
    if type_def is None:
        return {}

    if type_def.get("is_enum"):
        vals = type_def.get("enum_values", [])
        return vals[0] if vals else "example"

    result = {}
    for field in type_def.get("fields", []):
        fname = field["name"]
        ftype = field["type"]
        is_arr = field.get("is_array", False)

        val = _sample_value(ftype, types_dict, depth, seen)
        if is_arr:
            val = [val]
        result[fname] = val

    return result


def _sample_value(ftype, types_dict, depth, seen):
    """Return a sample value for a given type."""
    simple = {
        "string": "example",
        "boolean": False,
        "int": 0,
        "long": 0,
        "byte": 0,
        "double": 0.0,
        "float": 0.0,
        "dateTime": "2024-01-01T00:00:00",
        "timer": "00:00:00",
    }
    if ftype in simple:
        return simple[ftype]
    return build_sample_json(ftype, types_dict, depth + 1, seen)


def build_schema_help(type_name, types_dict, indent=0, seen=None):
    """Build formatted schema help text for a given type."""
    if seen is None:
        seen = set()
    if type_name in seen:
        return f"{'  ' * indent}(circular reference to {type_name})\n"
    seen = seen | {type_name}

    type_def = types_dict.get(type_name)
    if type_def is None:
        return ""

    lines = []
    if type_def.get("is_enum"):
        vals = ", ".join(type_def.get("enum_values", []))
        lines.append(f"{'  ' * indent}Enum: {vals}")
        return "\n".join(lines) + "\n"

    for field in type_def.get("fields", []):
        fname = field["name"]
        ftype = field["type"]
        is_arr = field.get("is_array", False)
        fdesc = field.get("description", "")

        type_display = ftype
        if is_arr:
            type_display = f"{ftype}[]"

        prefix = "  " * indent
        lines.append(f"{prefix}{fname} ({type_display}): {fdesc}")

        # Expand nested complex types one level deep (only at top level)
        nested_simple = {"string", "boolean", "int", "long", "double", "float", "dateTime", "timer"}
        if ftype not in nested_simple and indent < 1:
            nested = build_schema_help(ftype, types_dict, indent + 1, seen)
            if nested:
                lines.append(nested.rstrip())

    return "\n".join(lines) + "\n"


def is_complex_type(param_type):
    """Check if a parameter type is a complex type (not string/boolean/int/long)."""
    simple = {"string", "boolean", "int", "long", "byte", "double", "float", "dateTime", "timer"}
    base = param_type.split("[")[0].strip()
    return base not in simple


def is_array_type(param_type):
    """Check if a parameter type is an array."""
    return "[0..unbounded]" in param_type or "[]" in param_type


def param_flag_name(param_name):
    """Convert a camelCase param name to a kebab-case flag name."""
    result = re.sub(r"([a-z])([A-Z])", r"\1-\2", param_name).lower()
    return result


def go_flag_var(param_name):
    """Convert param name to Go variable for flag."""
    return "flag" + param_name[0].upper() + param_name[1:]


def generate_command_file(cli_cmd, methods, types_dict=None):
    """Generate a Go file for a CLI command group."""
    if types_dict is None:
        types_dict = {}
    var_name = to_go_var(cli_cmd)
    filename = to_go_filename(cli_cmd)

    # Descriptions for parent commands
    cmd_descriptions = {
        "agent-groups": "Manage agent groups",
        "call-variables": "Manage call variables and variable groups",
        "campaign-profiles": "Manage campaign profiles",
        "campaigns": "Manage campaigns (inbound, outbound, autodial)",
        "contacts": "Manage contacts and contact fields",
        "dialing-rules": "Manage dialing rules",
        "dispositions": "Manage dispositions",
        "domain-limits": "View domain limits",
        "ivr-scripts": "Manage IVR scripts",
        "lists": "Manage contact lists",
        "locales": "Manage locales and languages",
        "prompts": "Manage audio prompts",
        "reason-codes": "Manage reason codes",
        "reports": "Run and retrieve reports",
        "session": "Manage API sessions",
        "skills": "Manage skills",
        "speed-dial": "Manage speed dial numbers",
        "user-profiles": "Manage user profiles",
        "users": "Manage users",
        "vcc": "VCC configuration",
        "web-connectors": "Manage web connectors",
    }

    parent_desc = cmd_descriptions.get(cli_cmd, f"Manage {cli_cmd}")

    # Pre-scan: check if any method has complex params (need fmt import)
    has_complex = False
    for method_info in methods:
        for p in method_info.get("params", []):
            if is_complex_type(p["type"]):
                has_complex = True
                break
        if has_complex:
            break

    lines = []
    lines.append("// Code generated by codegen/generate_cli.py — DO NOT EDIT.")
    lines.append("//")
    lines.append("// To regenerate: python3 codegen/generate_cli.py")
    lines.append("")
    lines.append("package cmd")
    lines.append("")
    lines.append("import (")
    if has_complex:
        lines.append('\t"fmt"')
        lines.append("")
    lines.append('\t"github.com/Cloverhound/five9-cli/internal/output"')
    lines.append('\t"github.com/Cloverhound/five9-cli/internal/soap"')
    lines.append('\t"github.com/spf13/cobra"')
    lines.append(")")
    lines.append("")

    # Parent command
    lines.append(f"var {var_name}Cmd = &cobra.Command{{")
    lines.append(f'\tUse:   "{cli_cmd}",')
    lines.append(f'\tShort: "{parent_desc}",')
    lines.append("}")
    lines.append("")

    # Track subcommand names to avoid duplicates
    seen_subcmds = {}

    # Pre-generate sample JSON vars for commands with complex params
    sample_vars = {}  # subcmd_var -> sample var name
    for method_info in methods:
        soap_method = method_info["method"]
        subcmd_name = method_to_subcommand(soap_method, cli_cmd)
        params = method_info.get("params", [])

        if subcmd_name in seen_subcmds:
            continue

        complex_params = [p for p in params if is_complex_type(p["type"])]
        if complex_params and types_dict:
            cp = complex_params[0]
            base_type = cp["type"].split("[")[0].strip()
            sample = build_sample_json(base_type, types_dict)
            if sample:
                sample_var_name = to_go_var(cli_cmd) + to_go_var(subcmd_name) + "Sample"
                sample_json = json.dumps(sample, indent=2)
                lines.append(f"var {sample_var_name} = `{sample_json}`")
                lines.append("")
                sample_vars[subcmd_name] = sample_var_name

    seen_subcmds = {}

    # Generate subcommands
    for method_info in methods:
        soap_method = method_info["method"]
        subcmd_name = method_to_subcommand(soap_method, cli_cmd)
        params = method_info.get("params", [])
        desc = method_info.get("description", f"Call {soap_method}")
        # Escape quotes in description
        short_desc = desc.replace('"', '\\"')
        # Truncate long descriptions
        if len(short_desc) > 120:
            short_desc = short_desc[:117] + "..."

        # Handle duplicate subcommand names
        if subcmd_name in seen_subcmds:
            continue
        seen_subcmds[subcmd_name] = True

        subcmd_var = to_go_var(cli_cmd) + to_go_var(subcmd_name) + "Cmd"

        # Separate string params from complex params
        string_params = []
        complex_params = []
        array_string_params = []

        for p in params:
            ptype = p["type"]
            if is_array_type(ptype) and not is_complex_type(ptype):
                array_string_params.append(p)
            elif is_complex_type(ptype):
                complex_params.append(p)
            else:
                string_params.append(p)

        # Build Long description with schema help for complex params
        long_desc = None
        if complex_params and types_dict:
            cp = complex_params[0]
            base_type = cp["type"].split("[")[0].strip()
            schema = build_schema_help(base_type, types_dict)
            if schema:
                # Build long desc: original description + schema
                long_parts = []
                if desc:
                    long_parts.append(desc)
                long_parts.append("")
                long_parts.append(f"Body file schema ({base_type}):")
                long_parts.append(schema.rstrip())
                long_desc = "\n".join(long_parts)

        # Build the RunE function
        lines.append(f"var {subcmd_var} = &cobra.Command{{")
        lines.append(f'\tUse:   "{subcmd_name}",')
        lines.append(f'\tShort: "{short_desc}",')
        if long_desc:
            # Use Go raw string literal for Long
            escaped_long = long_desc.replace("`", "'")
            lines.append(f'\tLong:  `{escaped_long}`,')
        lines.append(f"\tRunE: func(cmd *cobra.Command, args []string) error {{")

        # Add --sample early return for commands with complex params
        if subcmd_name in sample_vars:
            lines.append(f'\t\tsample, _ := cmd.Flags().GetBool("sample")')
            lines.append(f"\t\tif sample {{")
            lines.append(f"\t\t\tfmt.Println({sample_vars[subcmd_name]})")
            lines.append(f"\t\t\treturn nil")
            lines.append(f"\t\t}}")

        lines.append(f'\t\treq := soap.NewRequest("{soap_method}")')

        # Add flag retrieval for string params
        for p in string_params:
            flag_name = param_flag_name(p["name"])
            var_name_local = go_flag_var(p["name"])
            ptype = p["type"]
            if ptype == "boolean":
                lines.append(f'\t\t{var_name_local}, _ := cmd.Flags().GetBool("{flag_name}")')
                lines.append(f'\t\tif {var_name_local} {{')
                lines.append(f'\t\t\treq.SetParam("{p["name"]}", "true")')
                lines.append(f"\t\t}}")
            else:
                lines.append(f'\t\t{var_name_local}, _ := cmd.Flags().GetString("{flag_name}")')
                lines.append(f'\t\treq.SetParam("{p["name"]}", {var_name_local})')

        # Add flag retrieval for array string params
        for p in array_string_params:
            flag_name = param_flag_name(p["name"])
            var_name_local = go_flag_var(p["name"])
            lines.append(f'\t\t{var_name_local}, _ := cmd.Flags().GetStringSlice("{flag_name}")')
            lines.append(f"\t\tfor _, v := range {var_name_local} {{")
            lines.append(f'\t\t\treq.SetParam("{p["name"]}", v)')
            lines.append(f"\t\t}}")

        # Handle complex params with --body-file
        if complex_params:
            wrapper = complex_params[0]["name"]
            lines.append(f'\t\tbodyFile, _ := cmd.Flags().GetString("body-file")')
            lines.append(f'\t\tif bodyFile != "" {{')
            lines.append(f'\t\t\tif err := req.SetBodyFile("{wrapper}", bodyFile); err != nil {{')
            lines.append(f"\t\t\t\treturn err")
            lines.append(f"\t\t\t}}")
            lines.append(f"\t\t}}")

        lines.append(f"\t\tdata, err := soap.Do(req)")
        lines.append(f"\t\tif err != nil {{")
        lines.append(f"\t\t\treturn err")
        lines.append(f"\t\t}}")
        lines.append(f"\t\treturn output.Print(data)")
        lines.append(f"\t}},")
        lines.append(f"}}")
        lines.append("")

    # init() function
    lines.append("func init() {")
    lines.append(f"\trootCmd.AddCommand({to_go_var(cli_cmd)}Cmd)")

    seen_init = set()
    for method_info in methods:
        soap_method = method_info["method"]
        subcmd_name = method_to_subcommand(soap_method, cli_cmd)
        params = method_info.get("params", [])

        # Skip duplicates
        if subcmd_name not in seen_subcmds or subcmd_name in seen_init:
            continue
        seen_init.add(subcmd_name)

        subcmd_var = to_go_var(cli_cmd) + to_go_var(subcmd_name) + "Cmd"
        lines.append(f"\t{to_go_var(cli_cmd)}Cmd.AddCommand({subcmd_var})")

        # Add flags
        string_params = []
        complex_params = []
        array_string_params = []

        for p in params:
            ptype = p["type"]
            if is_array_type(ptype) and not is_complex_type(ptype):
                array_string_params.append(p)
            elif is_complex_type(ptype):
                complex_params.append(p)
            else:
                string_params.append(p)

        for p in string_params:
            flag_name = param_flag_name(p["name"])
            pdesc = p.get("description", "").replace('"', '\\"')
            if len(pdesc) > 100:
                pdesc = pdesc[:97] + "..."
            ptype = p["type"]
            if ptype == "boolean":
                lines.append(f'\t{subcmd_var}.Flags().Bool("{flag_name}", false, "{pdesc}")')
            elif "pattern" in p["name"].lower():
                lines.append(f'\t{subcmd_var}.Flags().String("{flag_name}", ".*", "{pdesc}")')
            else:
                lines.append(f'\t{subcmd_var}.Flags().String("{flag_name}", "", "{pdesc}")')

        for p in array_string_params:
            flag_name = param_flag_name(p["name"])
            pdesc = p.get("description", "").replace('"', '\\"')
            if len(pdesc) > 100:
                pdesc = pdesc[:97] + "..."
            lines.append(f'\t{subcmd_var}.Flags().StringSlice("{flag_name}", nil, "{pdesc}")')

        if complex_params:
            wrapper = complex_params[0]["name"]
            pdesc = f"JSON file for {wrapper} parameter"
            lines.append(f'\t{subcmd_var}.Flags().String("body-file", "", "{pdesc}")')
            lines.append(f'\t{subcmd_var}.Flags().Bool("sample", false, "Print sample JSON body and exit")')

    lines.append("}")
    lines.append("")

    return filename, "\n".join(lines)


def main():
    with open(SPEC_PATH) as f:
        spec = json.load(f)

    commands = spec["commands"]
    types_dict = spec.get("types", {})

    # Skip reports — will be hand-written
    skip = {"reports"}

    generated = []
    for cli_cmd, methods in sorted(commands.items()):
        if cli_cmd in skip:
            continue

        filename, code = generate_command_file(cli_cmd, methods, types_dict)
        filepath = os.path.join(CMD_DIR, filename)

        with open(filepath, "w") as f:
            f.write(code)

        generated.append(filename)
        print(f"Generated {filename} ({len(methods)} methods)")

    # Generate reports with just the basic 4 SOAP methods (run-and-wait added manually)
    if "reports" in commands:
        filename, code = generate_command_file("reports", commands["reports"], types_dict)
        filepath = os.path.join(CMD_DIR, filename)
        with open(filepath, "w") as f:
            f.write(code)
        generated.append(filename)
        print(f"Generated {filename} ({len(commands['reports'])} methods)")

    print(f"\nGenerated {len(generated)} command files")


if __name__ == "__main__":
    main()
