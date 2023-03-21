#include <ctype.h>
#include <stdbool.h>
#include <string.h>
#include <stdio.h>
#include <stdlib.h>

typedef struct node {
	char src;
	char dst;
	struct node *next;
} node_t;

#define EDGE_LINE 1
#define QUESTION_LINE 2
#define UNKNOWN_LINE 0

int parse_linetype(const char* line) {
	if (strstr(line, "->")) return EDGE_LINE;
	if (strchr(line, ',')) return QUESTION_LINE;
	return UNKNOWN_LINE;
}

node_t* parse_edge(const char* line) {
	node_t *n = (node_t*)calloc(sizeof(node_t), 1);
	for(const char *p=line; *p; p++) {
		if (isalpha(*p)) {
			if (n->src == 0) {
				n->src = *p;
			} else if (n->dst == 0) {
				n->dst = *p;
			}
		}
	}
	return n;
}

/**
 * parse_question
 *
 * param const char *line
 * return char*
 *
 * Takes a comma separated list of characters, returns
 * the list as a null terminated array of characters.
 * 
 * The returns array is valid until parse_question is 
 * called again.
 */
char* parse_question(const char* line) {
	int idx=0;
	static char question[20];
	memset(question, 0, sizeof(question));

	for(const char *p=line; *p; p++) {
		if (isalpha(*p)) {
			question[idx++] = *p;
		}
	}
	return question;
}

char node_leads_to(node_t *haystack, char needle) {
	for(node_t *cur = haystack; cur; cur = cur->next) {
		if (cur->src == needle) {
			return cur->dst;
		}
	}
	return 0;
}

bool intersects(node_t* haystack, char a, char b) {
	char cura = a;
	char curb = b;

	while(cura) {
		curb = b;
		while(curb) {
			if (cura == curb) return true;
			curb = node_leads_to(haystack, curb);
		}
		cura = node_leads_to(haystack, cura);
	}
	return false;
}

bool set_intersects(node_t *haystack, const char* question) {

	for(int i=0; question[i+1]; ++i) {
		for(int j=i+1; question[j]; ++j) {
			if (intersects(haystack, question[i], question[j])) {
				return true;
			}
		}
	}

	return false;
}

int main(int argc, char** argv) {

	char line[80];
	char question[20];
	const char* intersect_label;
	node_t *EDGES = NULL;
	node_t *cur;
	while(fgets(line, sizeof(line), stdin)) {
		
		// Remove any line end characters
		for(char *p=line; *p; ++p) {
			switch(*p) {
			case '\n':
			case '\r':
				*p = 0;
				break;
			}
		}

		// Act on the lines
		switch(parse_linetype(line)) {
		case EDGE_LINE:
			cur = parse_edge(line);
			cur->next = EDGES;
			EDGES = cur;
			break;
		case QUESTION_LINE:
			if (set_intersects(EDGES, parse_question(line))) intersect_label = "true"; else intersect_label = "false";
			printf("%s: %s\n", line, intersect_label);
			break;
		default:
			// Do nothing	
			break;
		}
	}

	return 0;
}
