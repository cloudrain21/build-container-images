import os, sys
import configparser  # use ConfigParser for v2 (code must be modified)

def mk_docker_file(config):
    for section in config.sections():
        # read template file name
        template_file_name = config[section]["template"]
        template_file_name = "template/" + template_file_name

        # replace all "$tag" type tags to matching config values
        data = ""
        with open(template_file_name, "r") as f:
            # replace contents
            data = f.read()
            for key in config[section]:
                if key == "template":
                    continue
                val = config[section][key]
                src_substr = "\"$" + key + "\""
                data = data.replace(src_substr, val)

            # write replaced content to new dockerfile
            # filename = Dockerfile.section_name
            docker_file_name = "dockerfile/" + "Dockerfile" + "." + section
            with open(docker_file_name, "w+") as df:
                wlen = df.write(data)
                if wlen == len(data):
                    print("Done : %s" %(docker_file_name))
                else:
                    print("Fail : %s (%d) != (%d)" %(docker_file_name, wlen, len(data)))


def read_config(config_file):
    config = configparser.ConfigParser()
    config.read(config_file)
    return config

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage : %s config_file" %(os.path.basename(sys.argv[0])))
        sys.exit(1)

    config = read_config(sys.argv[1])
    mk_docker_file(config)

