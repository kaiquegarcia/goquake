package cmd

import "fmt"

func help() {
	fmt.Printf(`
###############################################
######## Welcome to goquake project ###########
###############################################
#### Author: Kaique Garcia Menezes ############
#### E-mail: contato@kaiquegarcia.dev #########
###############################################
#
# To use this, put a Quake Game log in
# the same folder of the program with 
# the filename 'qgames.log'.
#
# You can also customize some config:
===============================================
# '-f <txt>' or '--filepath <txt>'
# Use this argument to change the log filepath
# to be parsed. If it's not possible to find
# the file, an error will be printed as
# response.
===============================================
# '-wap <v>' or '--world-as-player <v>'
# Use this to consider <world> as a player in
# the kill calculations.
# Currently, the options to '<v>' are true or 1.
# Anything else will be interpreted as false.
===============================================
# '-gb <txt>' or '--group-by <txt>'
# Use this to change the 'group-by' behavior.
# Currently, the only option to '<txt>' is
# 'mean', so we group the result by 'mean'.
# Anything else will be ignored.
===============================================
# And that's all folks. Hope you enjoy it.
###############################################`)
}
