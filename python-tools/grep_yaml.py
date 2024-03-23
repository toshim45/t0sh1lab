import sys
import os
import yaml

if len(sys.argv) != 3:
	print("Grep yaml content value by key prefix/suffix, 1 deep inside directory")
	print("USAGE: python grep_yaml.py <dir> <filter-in-quote:*endswith,startswith*>")
	print("EXAMPLE: python grep_yaml.py /path/to/folder '*suffix_keyword'")
	sys.exit(1)

path_dir = sys.argv[1]
filter_word_content = sys.argv[2]

filter_ending = False
filter_beginning = False

if filter_word_content.startswith("*"):
	filter_ending = True
	filter_word_content = filter_word_content[1:]
elif filter_word_content.endswith("*"):
	filter_beginning = True
	filter_word_content = filter_word_content[0:-1]

if not os.path.exists(path_dir):
	print("ERROR: directory not exists %s" % path_dir)
	sys.exit(1)

result = set()

for dir_content in os.listdir(path_dir):
	if os.path.isfile(dir_content):
		continue
	
	# if filter_dir not in dir_content:
	# 	continue

	full_path_dir_content = (os.path.join(path_dir,dir_content))

	for sub_dir_content in os.listdir(full_path_dir_content):
		if os.path.isdir(sub_dir_content):
		    continue
		if not sub_dir_content.endswith('.yaml'):
			continue
		full_path_file_sub_dir_content = os.path.join(full_path_dir_content, sub_dir_content)
		with open(full_path_file_sub_dir_content,'r') as f:
			cfg = yaml.safe_load(f)
			cfgenvlist = cfg.get("env")
			if cfgenvlist is None:
				continue
			for cfgenv in cfgenvlist:
				if filter_beginning and not cfgenv.get("name").startswith(filter_word_content):
					continue
				if filter_ending and not cfgenv.get("name").endswith(filter_word_content):
					continue

				result.add("%s : %s " % (cfgenv.get("value"),sub_dir_content.replace("values.","")))


print(*result, sep="\n")