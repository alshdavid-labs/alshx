# Recursive Removal
## Deleting annoying files and folders

### Usage

The following command will look for `node_modules` and `.DS_Store` files (in order), recursively starting from `./` 

```bash
rrm --find node_modules --find .DS_Store ./
```

### Shorthand

```bash
# Same as:
# rrm --find node_modules ./

rrm -f node_modules
```