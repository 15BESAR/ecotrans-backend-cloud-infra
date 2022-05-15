# Author : David Fauzi 
# Desc : Edit deployment.yaml deployment name, replica, and image
import argparse
import yaml

if __name__ == '__main__':
    my_parser = argparse.ArgumentParser()
    my_parser.add_argument('yaml',type=str,help='yaml file')
    my_parser.add_argument('name',type=str,help='Deployment name')
    my_parser.add_argument('replica',type=str,help='replica number')
    my_parser.add_argument('project',type=str,help='project id')
    my_parser.add_argument('image',type=str,help='image')
    my_parser.add_argument('tag',type=str,help='tag')
    args = my_parser.parse_args()
    with open(args.yaml,'r') as f:
        data = yaml.safe_load(f)
        # edit deployment name
    data['metadata']['name']= args.name
    # edit replicas
    data['spec']['replicas']= int(args.replica)
    # edit image name: project + name + tag
    data['spec']['template']['spec']['containers'][0]['image'] = f'gcr.io/{args.project}/{args.image}:{args.tag}' 
    with open(args.yaml,'w') as f:
        yaml.dump(data,f)
    # print(yaml.dump(data, default_flow_style=False, sort_keys=False))
    print(f'Done updating {args.yaml} ....\n\n')