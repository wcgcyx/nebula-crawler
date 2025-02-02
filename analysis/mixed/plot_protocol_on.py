import psycopg2
import toml
import matplotlib.pyplot as plt
from lib import node_time, node_classification, node_protocol
import os
import psutil


config = toml.load("./db.toml")['psql']
conn = psycopg2.connect(
    host=config['host'],
    port=config['port'],
    database=config['database'],
    user=config['user'],
    password=config['password'],
)

start, end = node_time.get_time_range(conn)
all = node_classification.get_on_nodes(conn, start, end)
protocols = node_protocol.get_agent_protocol(conn, all)

counts = dict()
sum = 0
for _, protos in protocols.items():
    for proto in protos:
        if proto in counts:
            counts[proto] += 1
        else:
            counts[proto] = 1
        sum += 1
countsTrim = {"others": 0}
for key, val in counts.items():
    if val / sum < 0.01:
        countsTrim["others"] += val
    else:
        countsTrim[key] = val

# Plot
plt.rc('font', size=8)
plt.pie(countsTrim.values(), labels=countsTrim.keys(), autopct="%.1f%%")
plt.title("On nodes protocols from %s to %s" % (start.replace(microsecond=0), end.replace(microsecond=0)))
plt.savefig("./figs/protocols_for_on_nodes.png")
print("memory used:", psutil.Process(os.getpid()).memory_info().rss / 1024 ** 2, "MB")