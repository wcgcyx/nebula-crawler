import psycopg2
import toml
import matplotlib.pyplot as plt
from lib import node_time, node_classification, node_agent, node_cloud
import os
import psutil


# Helper function to trim agent version
def trim_agent(agent):
    if agent.startswith("/"):
        version = agent[1:]
    if agent.startswith("go-ipfs"):
        return "go-ipfs"
    elif agent.startswith("hydra-booster"):
        return "hydra-booster"
    elif agent.startswith("storm"):
        return "storm"
    elif agent.startswith("ioi"):
        return "ioi"
    else:
        return "others"


config = toml.load("./db.toml")['psql']
conn = psycopg2.connect(
    host=config['host'],
    port=config['port'],
    database=config['database'],
    user=config['user'],
    password=config['password'],
)

start, end = node_time.get_time_range(conn)
# Get storm node ids
all = node_classification.get_all_nodes(conn, start, end)
agents = node_agent.get_agent_version(conn, all)
storm = set()
for id, agent in agents.items():
    agent = trim_agent(agent)
    if agent == "storm":
        storm.add(id)

clouds = node_cloud.get_cloud(conn, storm)

counts = dict()
sum = 0
for _, cloud in clouds.items():
    if cloud in counts:
        counts[cloud] += 1
    else:
        counts[cloud] = 1

# Plot
plt.rc('font', size=8)
plt.pie(counts.values(), labels=counts.keys(), autopct="%.1f%%")
plt.title("Storm nodes cloud info from %s to %s" % (start.replace(microsecond=0), end.replace(microsecond=0)))
plt.savefig("./figs/cloud_info_for_all_storm.png")
print("memory used:", psutil.Process(os.getpid()).memory_info().rss / 1024 ** 2, "MB")